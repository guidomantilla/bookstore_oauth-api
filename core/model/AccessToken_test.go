package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(testing *testing.T) {

	accessToken := NewAccessToken(1)
	assert.NotNil(testing, accessToken, "accessToken must not be null")
	assert.False(testing, accessToken.IsExpired(), "accessToken must not be expired")
	assert.True(testing, accessToken.Token == "", "accessToken token should not have been defined")
	assert.True(testing, accessToken.UserId == 0, "accessToken userId should not have been defined")
	assert.True(testing, accessToken.DeviceId == 0, "accessToken deviceId should not have been defined")
	assert.True(testing, accessToken.AppId == 0, "accessToken appId should not have been defined")
}

func TestAccessToken_IsExpired(testing *testing.T) {
	accessToken := AccessToken{}
	assert.True(testing, accessToken.IsExpired(), "empty token must be expired")

	accessToken.ExpirationTime = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(testing, accessToken.IsExpired(), "token expiring three hours from now should not be expired")
}
