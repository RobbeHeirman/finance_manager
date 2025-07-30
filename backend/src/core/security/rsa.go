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
		return nil, errors.New("failed to decode PEM block private key")
	}

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

func ParseRsaPublicKey(pemString string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("failed to decode PEM block public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		slog.Error("Failed to parse public key", "err", err)
		return nil, errors.New("not RSA public key")
	}
	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		slog.Error("Failed to parse public key", "err", err)
		return nil, errors.New("not RSA public key")
	}
	return rsaPubKey, nil
}
