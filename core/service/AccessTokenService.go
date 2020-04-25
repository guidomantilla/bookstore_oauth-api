package service

import (
	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
)

type AccessTokenService interface {
	Create(accessToken *AccessToken) *Exception
	Update(id string, accessToken *AccessToken) *Exception
	Delete(id string) *Exception
	FindById(id string) (*AccessToken, *Exception)

	//
	UpdateExpirationTime(id string, expirationTime int64) *Exception
}
