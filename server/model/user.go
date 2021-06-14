package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type userInfo struct {
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(id int, name, username, email string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Username: username,
		Email:    email,
	}
}

func NewUserFromResponse(userResponse *http.Response) (*User, error) {
	user := User{}
	userBody, err := ioutil.ReadAll(userResponse.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(userBody, &user)

	return &user, nil
}
