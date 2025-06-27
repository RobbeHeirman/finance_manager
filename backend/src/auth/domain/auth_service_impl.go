package domain

import "errors"

type AuthServiceImpl struct {
	repository Repository
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
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

func (service *AuthServiceImpl) CreateJWT(user *User) (*string, error) {
	return nil, nil
}
