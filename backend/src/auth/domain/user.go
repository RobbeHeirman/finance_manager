package domain

import (
	"errors"
	"finance_manager/src/core/data_structures"
	"github.com/google/uuid"
)

type User struct {
	id        *uuid.UUID
	email     *data_structures.Email
	firstName *string
	lastName  *string
	imageURL  *data_structures.Url
}

func NewUser(id *uuid.UUID, email *data_structures.Email, firstName *string, lastName *string, imageURL *data_structures.Url) (*User, error) {
	if email.ToString() == "" {
		return nil, errors.New("email cannot be blank")
	}

	return &User{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		imageURL:  imageURL,
	}, nil
}

func (user *User) geId() data_structures.Optional[uuid.UUID] {
	return data_structures.CreateOptional(user.id)
}

func (user *User) geEmail() *data_structures.Email {
	return user.email
}

func (user *User) geFirstName() data_structures.Optional[string] {
	return data_structures.CreateOptional(user.firstName)
}

func (user *User) geLastName() data_structures.Optional[string] {
	return data_structures.CreateOptional(user.lastName)
}
func (user *User) geImageURL() data_structures.Optional[data_structures.Url] {
	return data_structures.CreateOptional(user.imageURL)
}
