package api

import (
	"errors"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type ResponseUser struct {
	User         User
	Applications []Application
}

type User struct {
	ID     string `json:"id"`
	Tag    string `json:"tag"`
	Locale string `json:"locale"`
	Email  string `json:"email"`
	Plan   struct {
		Name   string `json:"name"`
		Memory struct {
			Limit     int `json:"limit"`
			Available int `json:"available"`
			Used      int `json:"used"`
		} `json:"memory"`
		Duration struct {
			Formatted string      `json:"formatted"`
			Raw       interface{} `json:"raw"`
		} `json:"duration"`
	} `json:"plan"`
	Blocklist bool `json:"blocklist"`
}

func GetSelfUser() (*ResponseUser, error) {
	return GetUser("0")
}

func GetUser(id string) (*ResponseUser, error) {
	var path = EndpointUserSelfInfo
	if id != "0" {
		path = EndpointUserInfo(id)
	}

	res, err := request(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	if res.RawResponse.StatusCode != 200 {
		return nil, errors.New(ParseError(res))
	}

	var user *ResponseUser
	err = mapstructure.Decode(res.Response, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
