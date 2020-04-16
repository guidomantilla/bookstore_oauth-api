package repository

import (
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
)

type AccessTokenRepository interface {
	Create(accessToken *AccessToken) error
	Update(accessToken *AccessToken) error
	Delete(id string) error
	FindById(id string) (*AccessToken, error)
	Find(paramMap map[string][]string) (*[]AccessToken, error)
	ExistsById(id string) (bool, error)

	//
	UpdateExpirationTime(id string, expirationTime int64) error
}
