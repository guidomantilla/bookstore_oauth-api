package service

import (
	. "github.com/guidomantilla/bookstore_oauth-api/common/exception"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
)

type AccessTokenService interface {
	Create(accessToken *AccessToken) *Exception
	Update(accessToken *AccessToken) *Exception
	Delete(id string) *Exception
	FindById(id string) (*AccessToken, *Exception)
	Find(paramMap map[string][]string) (*[]AccessToken, *Exception)
	ExistsById(id string) (bool, *Exception)

	//
	UpdateExpirationTime(id string, expirationTime int64) *Exception
}
