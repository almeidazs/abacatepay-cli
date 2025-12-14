package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"time"
)

const baseURL = "https://api.abacatepay.com/v1"

type Client struct {
	key        string
	httpClient *http.Client
	maxRetries int
}

type APIResponse struct {
	Data  map[string]any `json:"data"`
	Error *string        `json:"error"`
}

func NewClient(key string) *Client {
	transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
	}

	return &Client{
		key:        key,
		maxRetries: 3,
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   15 * time.Second,
		},
	}
}

func backoff(attempt int) time.Duration {
	max := 2 * time.Second
	base := 200 * time.Millisecond

	duration := min(base * time.Duration(1<<attempt), max)

	jitter := time.Duration(rand.Int63n(int64(duration / 2)))

	return duration/2 + jitter
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		var netErr net.Error

		return errors.As(err, &netErr)
	}

	return resp != nil && (resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests)
}

func (c *Client) Get(route string) (map[string]any, error) {
	return c.Make(http.MethodGet, route, nil)
}

func (c *Client) Post(route string, body any) (map[string]any, error) {
	return c.Make(http.MethodPost, route, body)
}


func (c *Client) Make(
	method string,
	route string,
	body any,
) (map[string]any, error) {
	var err error
	var bodyBytes []byte

	if body != nil {
		bodyBytes, err = json.Marshal(body)
		
		if err != nil {
			return nil, err
		}
	}

	url := baseURL + route

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		var bodyReader io.Reader

		if bodyBytes != nil {
			bodyReader = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequest(method, url, bodyReader)
		if err != nil {
			return nil, err
		}

		if bodyBytes != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		req.Header.Set("Authorization", "Bearer "+c.key)

		resp, err := c.httpClient.Do(req)
		if !shouldRetry(err, resp) {
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				b, _ := io.ReadAll(resp.Body)
				return nil, fmt.Errorf("api error (%d): %s", resp.StatusCode, b)
			}

			var apiResp APIResponse
			if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
				return nil, err
			}

			if apiResp.Error != nil {
				return nil, errors.New(*apiResp.Error)
			}

			return apiResp.Data, nil
		}

		if attempt < c.maxRetries {
			time.Sleep(backoff(attempt))
			continue
		}

		if err != nil {
			return nil, err
		}

		return nil, errors.New("request failed after retries")
	}

	return nil, errors.New("unreachable")
}
