package data_structures

import "net/mail"

type Email struct {
	adres string
}

func NewEmail(adres string) (Email, error) {
	if _, err := mail.ParseAddress(adres); err != nil {
		return Email{}, err
	}
	return Email{adres: adres}, nil
}

func (email *Email) ToString() *string {
	return &email.adres
}
