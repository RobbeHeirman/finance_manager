package domain

type Repository interface {
	CreateUpdateUser(*User) (*User, error)
}
