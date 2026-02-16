package domain

import "errors"

var (
	ErrUserNameIsRequired     = errors.New("name is required")
	ErrUserEmailIsRequired    = errors.New("email is required")
	ErrUserEmailInvalid       = errors.New("email invalid")
	ErrUserPasswordIsRequired = errors.New("password is required")
	ErrUserPasswordInvalid    = errors.New("password invalid")
)

type User struct {
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user *User) error
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

	return user, nil
}
