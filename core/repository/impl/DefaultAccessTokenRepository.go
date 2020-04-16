package impl

import (
	. "github.com/guidomantilla/bookstore_oauth-api/common/nosql"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
)

const (
	STATEMENT_CREATE                 = "INSERT INTO access_tokens(\"token\", user_id, device_id, app_id, expiration_time) VALUES (?, ?, ?, ?, ?)"
	STATEMENT_UPDATE                 = "UPDATE access_tokens SET user_id = ?, device_id = ?, app_id = ?, expiration_time = ? WHERE \"token\" = ?"
	STATEMENT_UPDATE_EXPIRATION_TIME = "UPDATE access_tokens SET expiration_time = ? WHERE \"token\" = ?"
	STATEMENT_DELETE                 = ""
	STATEMENT_FIND_BY_ID             = "SELECT \"token\", user_id, device_id, app_id, expiration_time FROM access_tokens WHERE \"token\" = ?"
	STATEMENT_FIND                   = ""
	STATEMENT_EXISTS_BY_ID           = ""
)

type DefaultAccessTokenRepository struct {
	cassandraDataSource CassandraDataSource
}

func NewDefaultAccessTokenRepository(cassandraDataSource CassandraDataSource) *DefaultAccessTokenRepository {
	return &DefaultAccessTokenRepository{
		cassandraDataSource: cassandraDataSource,
	}
}

func (accessTokenRepository *DefaultAccessTokenRepository) Create(accessToken *AccessToken) error {

	session := accessTokenRepository.cassandraDataSource.GetSession()

	query := session.Query(STATEMENT_CREATE, accessToken.Token, accessToken.UserId, accessToken.DeviceId, accessToken.AppId, accessToken.ExpirationTime)
	if err := query.Exec(); err != nil {
		return err
	}

	return nil
}

func (accessTokenRepository *DefaultAccessTokenRepository) Update(accessToken *AccessToken) error {

	session := accessTokenRepository.cassandraDataSource.GetSession()

	query := session.Query(STATEMENT_UPDATE, accessToken.UserId, accessToken.DeviceId, accessToken.AppId, accessToken.ExpirationTime, accessToken.Token)

	if err := query.Exec(); err != nil {
		return err
	}

	return nil
}

func (accessTokenRepository *DefaultAccessTokenRepository) Delete(id string) error {
	return nil
}

func (accessTokenRepository *DefaultAccessTokenRepository) FindById(id string) (*AccessToken, error) {

	session := accessTokenRepository.cassandraDataSource.GetSession()

	query := session.Query(STATEMENT_FIND_BY_ID, id)

	var accessToken AccessToken
	err := query.Scan(&accessToken.Token, &accessToken.UserId, &accessToken.DeviceId, &accessToken.AppId, &accessToken.ExpirationTime)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (accessTokenRepository *DefaultAccessTokenRepository) Find(paramMap map[string][]string) (*[]AccessToken, error) {
	return nil, nil
}

func (accessTokenRepository *DefaultAccessTokenRepository) ExistsById(id string) (bool, error) {
	return false, nil
}

//
func (accessTokenRepository *DefaultAccessTokenRepository) UpdateExpirationTime(id string, expirationTime int64) error {

	session := accessTokenRepository.cassandraDataSource.GetSession()

	query := session.Query(STATEMENT_UPDATE_EXPIRATION_TIME, expirationTime, id)

	if err := query.Exec(); err != nil {
		return err
	}

	return nil
}
