package domain

import (
	"errors"
	"time"

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
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Save(user *User) error
	Update(user *User) error
	GetById(id string) (*User, error)
	Count() (int, error)
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, ErrUserNameIsRequired
	}

	if email == "" {
		return nil, ErrUserEmailIsRequired
	}

	if !isValidEmail(email) {
		return nil, ErrUserEmailInvalid
	}

	if password == "" {
		return nil, ErrUserPasswordIsRequired
	}

	if !isValidPassword(password) {
		return nil, ErrUserPasswordInvalid
	}

	now := time.Now()

	return &User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func Update(id, name, email string) (*User, error) {
	if id == "" {
		return nil, ErrUserIdIsRequired
	}

	if name == "" {
		return nil, ErrUserNameIsRequired
	}

	if email == "" {
		return nil, ErrUserEmailIsRequired
	}

	if !isValidEmail(email) {
		return nil, ErrUserEmailInvalid
	}

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		UpdatedAt: time.Now(),
	}, nil
}
