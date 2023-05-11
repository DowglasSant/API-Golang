package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Nick        string    `json:"nick,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("the name is required and cannot be blank")
	}

	if user.Nick == "" {
		return errors.New("the nick is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("the email is required and cannot be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("invalid email format")
	}

	if step == "register" && user.Password == "" {
		return errors.New("the password is required and cannot be blank")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHash)
	}

	return nil
}
