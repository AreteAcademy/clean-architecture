package user

import (
	"errors"

	"github.com/areteacademy/internal/domain"
	"gorm.io/gorm"
)

var ErrRepoUserIsNil = errors.New("user is nil")

type GormUserRepository struct {
	db *gorm.DB
}

func NewGoUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user *domain.User) error {
	if user == nil {
		return ErrRepoUserIsNil
	}

	model := ToRepository(user)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormUserRepository) Update(user *domain.User) error {
	if user == nil {
		return ErrRepoUserIsNil
	}

	model := ToRepository(user)

	result := r.db.
		Model(&UserGorm{}).
		Where("id = ?", user.ID).
		Updates(model)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *GormUserRepository) GetById(id string) (*domain.User, error) {
	var model UserGorm

	err := r.db.First(&model, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *GormUserRepository) Count() (int, error) {
	var count int64

	if err := r.db.Model(&UserGorm{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

var _ domain.UserRepository = (*GormUserRepository)(nil)
