//go:build dev

package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

//go:embed ../private.key
var privateKey string

//go:embed ../public.key
var publicKey string

func init() {
	slog.Info("Running  DEV Loading ENVS from .env and .key files")
	err := godotenv.Load()
	os.Setenv("RSA_PRIVATE_KEY", privateKey)
	os.Setenv("RSA_PUBLIC_KEY", publicKey)
	if err != nil {
		slog.Warn("Error loading .env file. No Env variables set from here.")
		return
	}
}
