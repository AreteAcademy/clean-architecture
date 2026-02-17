package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNameIsRequired     = errors.New("name is required")
	ErrUserEmailIsRequired    = errors.New("email is required")
	ErrUserEmailInvalid       = errors.New("email invalid")
	ErrUserPasswordIsRequired = errors.New("password is required")
	ErrUserPasswordInvalid    = errors.New("password invalid")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserIdIsRequired       = errors.New("id is required")
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user *User) error
	GetById(id string) (*User, error)
	Count() (int, error)
}

func NewUser(user *User) (*User, error) {
	if user.Name == "" {
		return nil, ErrUserNameIsRequired
	}

	if user.Email == "" {
		return nil, ErrUserEmailIsRequired
	}

	if !isValidEmail(user.Email) {
		return nil, ErrUserEmailInvalid
	}

	if user.Password == "" {
		return nil, ErrUserPasswordIsRequired
	}

	if !isValidPassword(user.Password) {
		return nil, ErrUserPasswordInvalid
	}

	return &User{
		ID:       uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
