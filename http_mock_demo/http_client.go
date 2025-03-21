package http_mock_demo

import (
	"encoding/json"
	"net/http"
)

func GetUserInfo(apiURL string) (*User, error) {
	resp, err := http.Get(apiURL + "/api/GetUserInfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

type User struct {
	UserId int    `json:"UserId"`
	Name   string `json:"Name"`
}
