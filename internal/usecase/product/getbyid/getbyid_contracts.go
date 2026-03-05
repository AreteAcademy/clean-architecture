package product

import "time"

type GetByIdProductInput struct {
	ID     string
	UserId string
}

type GetByIdProductOutput struct {
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
