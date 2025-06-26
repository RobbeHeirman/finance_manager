package domain

type AuthService interface {
	CreateUpdateUser(user *User) User
}
