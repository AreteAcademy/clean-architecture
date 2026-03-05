package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrProductIdIsRequired          = errors.New("id is required")
	ErrProductNotFound              = errors.New("product not found")
	ErrProductUserNotOwner          = errors.New("user not owner")
	ErrProductUserIdIsRequired      = errors.New("user id is required")
	ErrProductCategoryIdIsRequired  = errors.New("category id is required")
	ErrProductNameIsRequired        = errors.New("name is required")
	ErrProductDescriptionIsRequired = errors.New("description is required")
	ErrProductStatusIsRequired      = errors.New("status is required")
	ErrProductStatusInvalid         = errors.New("status invalid")
	ErrProductPriceInvalid          = errors.New("invalid price")
	ErrProductUserNotFound          = errors.New("user not found")
	ErrProductCategoryNotFound      = errors.New("category not found")
	ErrProductCategoryUserNotOwner  = errors.New("category user not owner")
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
	GetByIdAndUserId(id, userId string) (*Product, error)
	ListByUserId(userId string) ([]*Product, error)
	Count() (int, error)
}

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
)

func isValidProductStatus(status ProductStatus) bool {
	return status == ProductStatusActive || status == ProductStatusInactive
}

func NewProduct(
	userId,
	categoryId,
	name,
	description string,
	status ProductStatus,
	price int,
) (*Product, error) {
	if userId == "" {
		return nil, ErrProductUserIdIsRequired
	}

	if categoryId == "" {
		return nil, ErrProductCategoryIdIsRequired
	}

	if name == "" {
		return nil, ErrProductNameIsRequired
	}

	if description == "" {
		return nil, ErrProductDescriptionIsRequired
	}

	if status == "" {
		return nil, ErrProductStatusIsRequired
	}

	if !isValidProductStatus(status) {
		return nil, ErrProductStatusInvalid
	}

	if price <= 0 {
		return nil, ErrProductPriceInvalid
	}

	now := time.Now()

	return &Product{
		ID:          uuid.NewString(),
		UserId:      userId,
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
		Status:      string(status),
		Price:       price,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
