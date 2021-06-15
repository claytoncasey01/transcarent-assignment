package model

import (
	"encoding/json"
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

func NewUserFromResponse(userResponse []byte) (*User, error) {
	user := User{}
	json.Unmarshal(userResponse, &user)

	return &user, nil
}
