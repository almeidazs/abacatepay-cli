package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com.almeidazs/abacatepay-cli/internal/utils"
)

var (
	configPath string
	once       sync.Once
)

func getPath() (string, error) {
	var err error

	once.Do(func() {
		var home string

		home, err = os.UserHomeDir()

		if err != nil {
			return
		}

		configPath = filepath.Join(home, ".abacate", "abacate.json")
	})

	return configPath, err
}

func Sweep() error {
	path, err := getPath()

	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	return os.WriteFile(path, []byte("{}"), 0o600)
}

type Profile struct {
	CreatedAt string `json:"created_at"`
	Verified  bool   `json:"verified,omitempty"`
}

type Config struct {
	Profiles map[string]Profile `json:"profiles"`
	Current  string             `json:"current,omitempty"`
}

func Load() (*Config, error) {
	path, err := getPath()

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Profiles: make(map[string]Profile),
			}, nil
		}

		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Profiles == nil {
		cfg.Profiles = make(map[string]Profile)
	}

	return &cfg, nil
}

func (c *Config) Save(name, key string) error {
	if c.Exists(name) {
		return fmt.Errorf("a profile named \"%s\" already exists", name)
	}

	path, err := getPath()

	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	now := time.Now().Format(time.RFC3339)

	c.Current = name
	c.Profiles[name] = Profile{
		CreatedAt: now,
		Verified:  true,
	}

	if err := utils.SaveKeyring(name, key); err != nil {
		return err
	}

	if err := utils.SaveKeyring("current", key); err != nil {
		return err
	}

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o600)
}

func (c *Config) SetCurrent(name string) error {
	if c.Current == name {
		return fmt.Errorf("\"%s\" is alredy the current profile", name)
	}

	path, err := getPath()

	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	c.Current = name

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o600)
}

func (c *Config) Exists(name string) bool {
	_, ok := c.Profiles[name]
	return ok
}
