package category

import (
	"time"

	"github.com/areteacademy/internal/domain"
)

type CategoryGorm struct {
	ID        string    `gorm:"primaryKey"`
	UserId    string    `gorm:"index;not nul"`
	Name      string    `gorm:"not nul"`
	Status    string    `gorm:"not nul"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CategoryGorm) TableName() string {
	return "categories"
}

func (c *CategoryGorm) ToDomain() *domain.Category {
	return &domain.Category{
		ID:        c.ID,
		UserId:    c.UserId,
		Name:      c.Name,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToRepository(category *domain.Category) *CategoryGorm {
	return &CategoryGorm{
		ID:        category.ID,
		UserId:    category.UserId,
		Name:      category.Name,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
