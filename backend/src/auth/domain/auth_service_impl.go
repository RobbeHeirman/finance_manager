package domain

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (service *AuthServiceImpl) CreateUpdateUser(user *User) *User {
	return nil
}
