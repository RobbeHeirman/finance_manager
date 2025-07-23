package domain

import (
	"crypto/rsa"
	"github.com/google/uuid"
)

type ConfigRepo interface {
	GetPrivateKey() *rsa.PrivateKey
}

type Repository interface {
	CreateUpdateUser(*User) (*uuid.UUID, error)
}
