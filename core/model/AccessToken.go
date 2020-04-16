package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"crypto/sha256"
)

const (
	EXPIRATION_TIME = 24
)

type AccessToken struct {
	Token          string `json:"token"`
	UserId         int64  `json:"user_id"`
	DeviceId       int64  `json:"device_id"`
	AppId          int64  `json:"app_id"`
	ExpirationTime int64  `json:"expiration_time"`
}

func NewAccessToken(userId int64) *AccessToken {

	expirationTime := time.Now().UTC().Add(EXPIRATION_TIME * time.Hour).Unix()

	plainToken := fmt.Sprintf("at-%d-%d-ran", userId, expirationTime)
	sha256Token := sha256.Sum256([]byte(plainToken))
	token := fmt.Sprintf("%x", sha256Token)

	return &AccessToken{
		Token:          token,
		UserId:         userId,
		ExpirationTime: expirationTime,
	}
}

func (accessToken AccessToken) IsExpired() bool {

	now := time.Now().UTC()
	expirationTime := time.Unix(accessToken.ExpirationTime, 0)

	return now.After(expirationTime)
}

func (accessToken AccessToken) Validate() error {
	accessToken.Token = strings.TrimSpace(accessToken.Token)
	if accessToken.Token == "" {
		return errors.New("access token has an invalid id")
	}

	if accessToken.UserId <= 0 {
		return errors.New("access token has an user id")
	}

	if accessToken.DeviceId <= 0 {
		return errors.New("access token has an device id")
	}

	if accessToken.AppId <= 0 {
		return errors.New("access token has an app id")
	}

	if accessToken.ExpirationTime <= 0 {
		return errors.New("access token has an expiration time")
	}

	return nil
}
