package domain

import "github.com/google/uuid"

type Repository interface {
	CreateUpdateUser(*User) (*uuid.UUID, error)
}
