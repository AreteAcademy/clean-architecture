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

func validEntity(
	userId,
	categoryId,
	name,
	description string,
	status ProductStatus,
	price int,
) error {
	if userId == "" {
		return ErrProductUserIdIsRequired
	}

	if categoryId == "" {
		return ErrProductCategoryIdIsRequired
	}

	if name == "" {
		return ErrProductNameIsRequired
	}

	if description == "" {
		return ErrProductDescriptionIsRequired
	}

	if status == "" {
		return ErrProductStatusIsRequired
	}

	if !isValidProductStatus(status) {
		return ErrProductStatusInvalid
	}

	if price <= 0 {
		return ErrProductPriceInvalid
	}

	return nil
}

func NewProduct(
	userId,
	categoryId,
	name,
	description string,
	status ProductStatus,
	price int,
) (*Product, error) {
	err := validEntity(
		userId,
		categoryId,
		name,
		description,
		status,
		price,
	)
	if err != nil {
		return nil, err
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

func (p *Product) UpdateProduct(
	categoryId,
	name,
	description string,
	status ProductStatus,
	price int,
) error {
	err := validEntity(
		p.UserId,
		categoryId,
		name,
		description,
		status,
		price,
	)
	if err != nil {
		return err
	}

	now := time.Now()

	p.CategoryId = categoryId
	p.Name = name
	p.Description = description
	p.Status = string(status)
	p.Price = price
	p.UpdatedAt = now

	return nil
}
