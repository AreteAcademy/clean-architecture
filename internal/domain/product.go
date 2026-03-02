package domain

import (
	"errors"
	"time"
)

var (
	ErrProductCategoryIdIsRequired = errors.New("category id is required")
)

type Product struct {
	ID          string
	UserId      string
	CategoryId  string
	Name        string
	Price       int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRepository interface {
	Save(product *Product) error
	Update(product *Product) error
	GetById(id string) (*Product, error)
	ListByUserId(userId string) ([]*Product, error)
	Count() (int, error)
}
