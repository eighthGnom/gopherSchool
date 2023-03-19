package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Password           string `json:"password,omitempty"`
	EnscriptedPassword string `json:"-"`
}

func (user *User) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.By(validateIf(user.EnscriptedPassword == "")), validation.Length(6, 100)))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) EnscriptPassword() error {
	hashedPassword, err := enscriptString(user.Password)
	if err != nil {
		return err
	}
	user.EnscriptedPassword = hashedPassword
	return nil
}
func (user *User) CompareEnscriptedPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EnscriptedPassword), []byte(password)) == nil
}

func enscriptString(str string) (string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordByte), nil
}

func (user *User) Sanitize() {
	user.Password = ""
}
