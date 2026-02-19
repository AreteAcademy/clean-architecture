package user

import "time"

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
