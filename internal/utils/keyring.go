package utils

import (
	"github.com/zalando/go-keyring"
)

const service = "abacate"

func SweepKeyrings() error {
	return keyring.DeleteAll(service)
}

func SaveKeyring(name, key string) error {
	return keyring.Set(service, name, key)
}

func GetKeyring(name string) (string, error) {
	return keyring.Get(service, name)
}

func DelKeyring(name string) error {
	return keyring.Delete(service, name)
}
