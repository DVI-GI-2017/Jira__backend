package models

import (
	"errors"

	"regexp"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrEmptyEmail       = errors.New("empty email")
	ErrWrongEmailFormat = errors.New("wrong email format")
)

type Email string

//Long and strange regexp to validate email format.
var emailRegex = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@“]+(\.[^<>()\[\]\\.,;:\s@“]+)*)|(“.+“))@((\[[0-9]{1,3}\.
[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

// Returns validation error if invalid or nil if valid
func (e Email) Validate() error {
	if len(e) == 0 {
		return ErrEmptyEmail
	}
	if !emailRegex.MatchString(string(e)) {
		return ErrWrongEmailFormat
	}
	return nil
}

var (
	ErrEmptyPassword       = errors.New("empty password")
	ErrWrongPasswordFormat = errors.New("wrong password format")
)

type Password string

var passwordRegex = regexp.MustCompile(`^[0-9a-zA-Z\s\r\n@!#$^%&*()+=\-\[\]\\';,./{}|":<>?]{3,14}$`)

// Validates passowrd
func (p Password) Validate() error {
	if len(p) == 0 {
		return ErrEmptyPassword
	}
	if !passwordRegex.MatchString(string(p)) {
		return ErrWrongPasswordFormat
	}
	return nil
}

var (
	ErrEmptyName       = errors.New("empty name")
	ErrWrongNameFormat = errors.New("wrong name format")
)

type Name string

var nameRegex = regexp.MustCompile(`^[a-zA-Z].{1,49}$`)

// Validates names
func (n Name) Validate() error {
	if len(n) == 0 {
		return ErrEmptyName
	}
	if !nameRegex.MatchString(string(n)) {
		return ErrWrongNameFormat
	}
	return nil
}

type User struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email    Email         `json:"email" bson:"email"`
	Password Password      `json:"password" bson:"password"`
	Name     Name          `json:"name" bson:"name"`
	Bio      string        `json:"bio" bson:"bio,omitempty"`
}

// Returns validation error or nil if valid
func (u User) Validate() error {
	if err := u.Email.Validate(); err != nil {
		return err
	}
	if err := u.Password.Validate(); err != nil {
		return err
	}
	if err := u.Name.Validate(); err != nil {
		return err
	}
	return nil
}

type UsersList []User
