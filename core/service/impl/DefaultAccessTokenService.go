package impl

import (
	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
	. "github.com/guidomantilla/bookstore_oauth-api/core/repository"
)

const (
	CREATE_ERROR_TITLE     = "error creating the access token"
	UPDATE_ERROR_TITLE     = "error updating the access token"
	DELETE_ERROR_TITLE     = "error deleting the access token"
	FIND_BY_ID_ERROR_TITLE = "error finding the access token"
	FIND_ERROR_TITLE       = "error finding the access tokens"
)

type DefaultAccessTokenService struct {
	accessTokenRepository AccessTokenRepository
}

func NewDefaultAccessTokenService(accessTokenRepository AccessTokenRepository) *DefaultAccessTokenService {
	return &DefaultAccessTokenService{
		accessTokenRepository: accessTokenRepository,
	}
}

func (accessTokenService *DefaultAccessTokenService) Create(accessToken *AccessToken) *Exception {

	return nil
}

func (accessTokenService *DefaultAccessTokenService) Update(accessToken *AccessToken) *Exception {
	return nil
}

func (accessTokenService *DefaultAccessTokenService) Delete(id string) *Exception {
	return nil
}

func (accessTokenService *DefaultAccessTokenService) FindById(id string) (*AccessToken, *Exception) {
	return nil, nil
}

func (accessTokenService *DefaultAccessTokenService) Find(paramMap map[string][]string) (*[]AccessToken, *Exception) {
	return nil, nil
}

func (accessTokenService *DefaultAccessTokenService) ExistsById(id string) (bool, *Exception) {
	return false, nil
}

//
func (accessTokenService *DefaultAccessTokenService) UpdateExpirationTime(id string, expirationTime int64) *Exception {
	return nil
}
