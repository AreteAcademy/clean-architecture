package category

import "time"

type CreateCategoryInput struct {
	UserId string
	Name   string
	Status string
}

type CreateCategoryOutput struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
