package product

import "time"

type CreateProductInput struct {
	UserId      string
	CategoryId  string
	Name        string
	Description string
	Status      string
	Price       int
}

type CreateProductOutput struct {
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
