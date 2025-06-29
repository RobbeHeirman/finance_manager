package domain

type AuthService interface {
	CreateUpdateUser(user *User) (*User, error)
	CreateJWTToken(user *User) (string, error)
}
