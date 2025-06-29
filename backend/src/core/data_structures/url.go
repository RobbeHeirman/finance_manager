package data_structures

import url2 "net/url"

type Url struct {
	url string
}

func CreateUrl(url string) (Url, error) {
	if _, err := url2.ParseRequestURI(url); err != nil {
		return Url{}, err
	}
	return Url{url}, nil
}

func (u *Url) ToString() *string {
	return &u.url
}
