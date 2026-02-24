package category

import "time"

type CategoryItem struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListUserByIdCategoryOutput []CategoryItem
