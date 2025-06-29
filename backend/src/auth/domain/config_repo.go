package domain

type ConfigRepo interface {
	GetPrivateKey() string
}
