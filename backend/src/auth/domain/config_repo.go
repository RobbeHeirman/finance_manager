package domain

import "crypto/rsa"

type ConfigRepo interface {
	GetPrivateKey() *rsa.PrivateKey
}
