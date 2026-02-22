package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCategoryUserIdIsRequired = errors.New("user id is required")
	ErrCategoryNameIsRequired   = errors.New("name is required")
	ErrCategoryStatusIsRequired = errors.New("status is required")
	ErrCategoryStatusInvalid    = errors.New("status invalid")
	ErrCategoryUserNotFound     = errors.New("user not found")
)

type CategoryStatus string

type Category struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	CategoryStatusActive   CategoryStatus = "ACTIVE"
	CategoryStatusInactive CategoryStatus = "INACTIVE"
)

type CategoryRepository interface {
	Save(category *Category) error
	Update(category *Category) error
	GetById(id string) (*Category, error)
	Count() (int, error)
}

func isValidCategoryStatus(status CategoryStatus) bool {
	return status == CategoryStatusActive || status == CategoryStatusInactive
}

func NewCategory(userId string, name string, status CategoryStatus) (*Category, error) {
	if userId == "" {
		return nil, ErrCategoryUserIdIsRequired
	}

	if name == "" {
		return nil, ErrCategoryNameIsRequired
	}

	if status == "" {
		return nil, ErrCategoryStatusIsRequired
	}

	if !isValidCategoryStatus(status) {
		return nil, ErrCategoryStatusInvalid
	}

	now := time.Now()

	return &Category{
		ID:        uuid.NewString(),
		UserId:    userId,
		Name:      name,
		Status:    string(status),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
