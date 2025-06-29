package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log/slog"
)

func ParseRSAPrivateKey(pemString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	// Try PKCS1 format first
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		slog.Error("Failed to parse private key", "err", err)
		return nil, err
	}
	privateKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		slog.Error("Failed to parse private key", "err", err)
		return nil, errors.New("not RSA private key")
	}

	return privateKey, nil
}
