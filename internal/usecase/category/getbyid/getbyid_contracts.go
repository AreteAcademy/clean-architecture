package category

import "time"

type GetByIdCategoryInput struct {
	ID     string
	UserId string
}

type GetByIdCategoryOutput struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
