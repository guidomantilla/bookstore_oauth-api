package impl

import (
	"errors"

	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
	. "github.com/guidomantilla/bookstore_oauth-api/core/repository"
)

const (
	CREATE_ERROR_TITLE     = "error creating the access token"
	UPDATE_ERROR_TITLE     = "error updating the access token"
	DELETE_ERROR_TITLE     = "error deleting the access token"
	FIND_BY_ID_ERROR_TITLE = "error finding the access token"
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

	//if err := Validate(user); err != nil {
	//	return BadRequestException(UPDATE_ERROR_TITLE, err)
	//}

	exists, err := accessTokenService.accessTokenRepository.ExistsById(accessToken.Token)
	if err != nil {
		return InternalServerErrorException(CREATE_ERROR_TITLE, err)
	}

	if exists {
		return NotFoundException(CREATE_ERROR_TITLE, errors.New("access token does exist"))
	}

	if err := accessTokenService.accessTokenRepository.Create(accessToken); err != nil {
		return InternalServerErrorException(CREATE_ERROR_TITLE, err)
	}

	return nil
}

func (accessTokenService *DefaultAccessTokenService) Update(id string, accessToken *AccessToken) *Exception {

	//if err := Validate(user); err != nil {
	//	return BadRequestException(UPDATE_ERROR_TITLE, err)
	//}

	exists, err := accessTokenService.accessTokenRepository.ExistsById(accessToken.Token)
	if err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	if !exists {
		return NotFoundException(UPDATE_ERROR_TITLE, errors.New("access token does not exist"))
	}

	accessToken.Token = id
	if err := accessTokenService.accessTokenRepository.Update(accessToken); err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	return nil
}

func (accessTokenService *DefaultAccessTokenService) Delete(id string) *Exception {

	exists, err := accessTokenService.accessTokenRepository.ExistsById(id)
	if err != nil {
		return InternalServerErrorException(DELETE_ERROR_TITLE, err)
	}

	if !exists {
		return NotFoundException(DELETE_ERROR_TITLE, errors.New("access token does not exist"))
	}

	if err := accessTokenService.accessTokenRepository.Delete(id); err != nil {
		return InternalServerErrorException(DELETE_ERROR_TITLE, err)
	}

	return nil
}

func (accessTokenService *DefaultAccessTokenService) FindById(id string) (*AccessToken, *Exception) {

	accessToken, err := accessTokenService.accessTokenRepository.FindById(id)
	if err != nil {

		if err.Error() == "not found" {
			return nil, NotFoundException(FIND_BY_ID_ERROR_TITLE, errors.New("access token does not exist"))
		}

		return nil, InternalServerErrorException(FIND_BY_ID_ERROR_TITLE, err)
	}

	return accessToken, nil
}

//
func (accessTokenService *DefaultAccessTokenService) UpdateExpirationTime(id string, expirationTime int64) *Exception {

	//if err := Validate(user); err != nil {
	//	return BadRequestException(UPDATE_ERROR_TITLE, err)
	//}

	exists, err := accessTokenService.accessTokenRepository.ExistsById(id)
	if err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	if !exists {
		return NotFoundException(UPDATE_ERROR_TITLE, errors.New("access token does not exist"))
	}

	if err := accessTokenService.accessTokenRepository.UpdateExpirationTime(id, expirationTime); err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	return nil
}
