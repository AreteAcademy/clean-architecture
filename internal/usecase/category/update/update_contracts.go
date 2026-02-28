package category

import "time"

type UpdateCategoryInput struct {
	ID     string
	UserId string
	Name   string
	Status string
}

type UpdateCategoryOutput struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
