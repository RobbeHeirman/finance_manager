package domain

type AuthServiceImpl struct {
	repository Repository
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (service *AuthServiceImpl) CreateUpdateUser(user *User) (*User, error) {
	return service.repository.CreateUpdateUser(user)
}
