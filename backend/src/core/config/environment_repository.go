package config

import "os"

type EnvironmentRepository struct {
}

func NewEnvironmentRepository() *EnvironmentRepository {
	return &EnvironmentRepository{}
}

func (repo *EnvironmentRepository) GetPrivateKey() string {
	return os.Getenv("RSA_PRIVATE_KEY")
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
	return os.Getenv("DB_USERNAME")
}

func (repo *EnvironmentRepository) GetDatabasePassword() string {
	return os.Getenv("DB_PASSWORD")
}
