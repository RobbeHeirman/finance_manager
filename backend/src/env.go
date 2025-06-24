//go:build dev

package main

import (
	"github.com/joho/godotenv"
	"log/slog"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file. No Env variables set from here.")
		return
	}
}
