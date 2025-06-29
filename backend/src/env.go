//go:build dev

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func init() {
	slog.Info("Running  DEV Loading ENVS from .env and .key files")
	err := godotenv.Load()

	privateKeyPath := "../private.key" // or absolute path
	publicKeyPath := "../public.key"

	privateKey, err := loadKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	publicKey, err := loadKey(publicKeyPath)
	if err != nil {
		panic(err)
	}

	os.Setenv("RSA_PRIVATE_KEY", privateKey)
	os.Setenv("RSA_PUBLIC_KEY", publicKey)
	if err != nil {
		slog.Warn("Error loading .env file. No Env variables set from here.")
		return
	}
}

func loadKey(path string) (string, error) {
	data, err := os.ReadFile(path) // Go 1.16+
	if err != nil {
		return "", fmt.Errorf("failed to read key %q: %w", path, err)
	}
	return string(data), nil
}
