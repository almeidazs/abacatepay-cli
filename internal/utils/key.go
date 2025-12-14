package utils

import (
	"errors"

	"github.com.almeidazs/abacatepay-cli/internal/prompts"
	"github.com/zalando/go-keyring"
)

func GetAPIKey() (string, error) {
	key, err := GetKeyring("current")

	if err == nil {
		return key, nil
	}

	if !errors.Is(err, keyring.ErrNotFound) {
		return "", err
	}

	return prompts.AskAPIKey()
}
