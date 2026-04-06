package product

import (
	"time"

	"github.com/areteacademy/internal/domain"
)

type ProductGorm struct {
	ID          string    `gorm:"primaryKey"`
	UserId      string    `gorm:"index;not nul"`
	CategoryId  string    `gorm:"index;not nul"`
	Name        string    `gorm:"not nul"`
	Description string    `gorm:"not nul"`
	Status      string    `gorm:"index;not nul"`
	Price       int       `gorm:"not nul"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (ProductGorm) TableName() string {
	return "products"
}

func (u *ProductGorm) ToDomain() *domain.Product {
	return &domain.Product{
		ID:          u.ID,
		UserId:      u.UserId,
		CategoryId:  u.CategoryId,
		Name:        u.Name,
		Description: u.Description,
		Status:      u.Status,
		Price:       u.Price,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func ToRepository(u *domain.Product) *ProductGorm {
	return &ProductGorm{
		ID:          u.ID,
		UserId:      u.UserId,
		CategoryId:  u.CategoryId,
		Name:        u.Name,
		Description: u.Description,
		Status:      u.Status,
		Price:       u.Price,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
