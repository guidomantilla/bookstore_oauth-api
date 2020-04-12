package model

import (
	"time"
)

const (
	EXPIRATION_TIME = 24
)

type AccessToken struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	DeviceId int64  `json:"device_id"`
	AppId    int64  `json:"app_id"`
	Expires  int64  `json:"expires"`
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME * time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {

	now := time.Now().UTC()
	expirationTime := time.Unix(accessToken.Expires, 0)

	return now.After(expirationTime)
}
