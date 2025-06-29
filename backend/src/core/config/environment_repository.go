package config

import (
	"crypto/rsa"
	"finance_manager/src/core/security"
	"log/slog"
	"os"
)

type EnvironmentRepository struct {
	privateKey *rsa.PrivateKey
}

func NewEnvironmentRepository() *EnvironmentRepository {

	return &EnvironmentRepository{}
}

func (repo *EnvironmentRepository) GetPrivateKey() *rsa.PrivateKey {
	if repo.privateKey == nil {
		key, err := security.ParseRSAPrivateKey(os.Getenv("RSA_PRIVATE_KEY"))
		if err != nil {
			slog.Error("Error parsing RSA_PRIVATE_KEY", "err", err)
		}
		repo.privateKey = key
	}
	return repo.privateKey
}

func (repo *EnvironmentRepository) GetPublicKey() string {
	return os.Getenv("RSA_PUBLIC_KEY")
}

func (repo *EnvironmentRepository) GetDatabaseHost() string {
	return os.Getenv("DB_HOST")
}

func (repo *EnvironmentRepository) GetDatabasePort() string {
	return os.Getenv("DB_PORT")
}

func (repo *EnvironmentRepository) GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}

func (repo *EnvironmentRepository) GetDatabaseUser() string {
	return os.Getenv("DB_USER")
}

func (repo *EnvironmentRepository) GetDatabasePassword() string {
	return os.Getenv("DB_PASSWORD")
}
