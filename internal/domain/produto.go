package domain

import (
	"errors"
	"time"
)

var (
	ErrProductUserIdIsRequired = errors.New("user id is required")
)

type Product struct {
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

type ProductRepository interface {
	Save(product *Product) error
	Update(product *Product) error
	GetById(id string) (*Product, error)
	ListByUserId(userId string) ([]*Product, error)
	Count() (int, error)
}
