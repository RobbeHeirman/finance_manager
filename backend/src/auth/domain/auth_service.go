package domain

import (
	"errors"
	"finance_manager/src/core/security"
)

type AuthService interface {
	CreateUpdateUser(user *User) (*User, error)
	CreateJWTToken(user *User) (string, error)
}

type AuthServiceImpl struct {
	repository Repository
	configRepo ConfigRepo
}

func NewAuthServiceImpl(config ConfigRepo, repository Repository) *AuthServiceImpl {
	return &AuthServiceImpl{
		repository: repository,
		configRepo: config,
	}
}

// CreateUpdateUser Creates or updates a user based on their email in the persistence layer.
// If the email address does not exist. Create a new user and return the new uuid.
// If it does exist. Update the other fields based on user in the persistence layer and return the existing uuid.
// Will set the uuid of the user based on what the persistence layer returns
// User.email cannot be nil.
func (service *AuthServiceImpl) CreateUpdateUser(user *User) (*User, error) {
	if user.GetEmail() == nil {
		return nil, errors.New("No email was set")
	}

	id, err := service.repository.CreateUpdateUser(user)
	if err != nil {
		return nil, err
	}
	user.SetId(id)
	return user, nil
}

func (service *AuthServiceImpl) CreateJWTToken(user *User) (string, error) {
	privateKey := service.configRepo.GetPrivateKey()
	return security.CreateJWT(privateKey, user.id.String())
}
