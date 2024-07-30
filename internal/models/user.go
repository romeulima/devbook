package models

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(userRequest *UserRequest) *User {
	return &User{
		Name:     userRequest.Name,
		Nick:     userRequest.Nick,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}

type UserRequest struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) ValidadeFields(userRequest *UserRequest) error {
	if userRequest.Name == "" {
		return errors.New("the field [name] is required")
	}

	if userRequest.Nick == "" {
		return errors.New("the field [nick] is required")
	}

	if userRequest.Email == "" {
		return errors.New("the field [email] is required")
	}

	if userRequest.Password == "" {
		return errors.New("the field [password] is required")
	}

	return nil
}
