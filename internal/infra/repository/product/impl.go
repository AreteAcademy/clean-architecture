package product

import (
	"errors"

	"github.com/areteacademy/internal/domain"
	"gorm.io/gorm"
)

var ErrRepoProductIsNil = errors.New("product is nil")

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) Save(product *domain.Product) error {
	if product == nil {
		return ErrRepoProductIsNil
	}

	model := ToRepository(product)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormProductRepository) Update(product *domain.Product) error {
	if product == nil {
		return ErrRepoProductIsNil
	}

	model := ToRepository(product)

	result := r.db.
		Model(&ProductGorm{}).
		Where("id = ?", product.ID).
		Updates(model)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

func (r *GormProductRepository) GetById(id string) (*domain.Product, error) {
	var models ProductGorm

	err := r.db.First(&models, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return models.ToDomain(), nil
}

func (r *GormProductRepository) GetByIdAndUserId(id, userId string) (*domain.Product, error) {
	var models ProductGorm

	err := r.db.First(&models, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return models.ToDomain(), nil
}

func (r *GormProductRepository) ListByUserId(userId string) ([]*domain.Product, error) {
	var models []ProductGorm

	err := r.db.Find(&models, "user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	products := make([]*domain.Product, 0, len(models))
	for _, model := range models {
		products = append(products, model.ToDomain())
	}

	return products, nil
}

func (r *GormProductRepository) Count() (int, error) {
	var count int64

	if err := r.db.Model(&ProductGorm{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

var _ domain.ProductRepository = (*GormProductRepository)(nil)
