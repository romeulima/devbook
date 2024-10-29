package models

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
	"github.com/romeulima/devbook/internal/security"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
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

func (u *UserRequest) Prepare(stage string, userRequest *UserRequest) error {
	if err := u.ValidadeFields(stage, userRequest); err != nil {
		return err
	}

	encryptedPassword, err := security.Encrypt(userRequest.Password)

	if err != nil {
		return err
	}

	userRequest.Password = string(encryptedPassword)

	return nil
}

func (u *UserRequest) ValidadeFields(stage string, userRequest *UserRequest) error {
	if userRequest.Name == "" {
		return errors.New("the field [name] is required")
	}

	if userRequest.Nick == "" {
		return errors.New("the field [nick] is required")
	}

	if userRequest.Email == "" {
		return errors.New("the field [email] is required")
	}

	if err := checkmail.ValidateFormat(userRequest.Email); err != nil {
		return errors.New("this email is invalid")
	}

	if stage == "cadastro" && userRequest.Password == "" {
		return errors.New("the field [password] is required")
	}

	return nil
}
