package category

import (
	"errors"

	"github.com/areteacademy/internal/domain"
	"gorm.io/gorm"
)

var ErrRepositoryCategoryNil = errors.New("category is nil")

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (r *GormCategoryRepository) Save(category *domain.Category) error {
	if category == nil {
		return ErrRepositoryCategoryNil
	}

	model := ToRepository(category)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormCategoryRepository) Update(category *domain.Category) error {
	if category == nil {
		return ErrRepositoryCategoryNil
	}

	model := ToRepository(category)

	result := r.db.
		Model(&CategoryGorm{}).
		Where("id = ?", category.ID).
		Updates(model)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}

func (r *GormCategoryRepository) GetById(id string) (*domain.Category, error) {
	var model CategoryGorm

	err := r.db.First(&model, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}

		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *GormCategoryRepository) GetByIdAndUserId(id, userId string) (*domain.Category, error) {
	var model CategoryGorm

	err := r.db.First(&model, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}

		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *GormCategoryRepository) ListByUserId(userId string) ([]*domain.Category, error) {
	var models []CategoryGorm

	if err := r.db.Where("user_id = ?", userId).Find(&models).Error; err != nil {
		return nil, err
	}

	categories := make([]*domain.Category, 0, len(models))
	for i := range models {
		categories = append(categories, models[i].ToDomain())
	}

	return categories, nil
}

func (r *GormCategoryRepository) Count() (int, error) {
	var count int64

	if err := r.db.Model(&CategoryGorm{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

var _ domain.CategoryRepository = (*GormCategoryRepository)(nil)
