package product

import "time"

type UpdateProductInput struct {
	ID          string
	CategoryId  string
	Name        string
	Description string
	Status      string
	Price       int
}

type UpdateProductOutput struct {
	ID          string
	UserId      string
	CategoryId  string
	Name        string
	Description string
	Status      string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
