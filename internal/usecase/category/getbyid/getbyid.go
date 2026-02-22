package category

import (
	"github.com/areteacademy/internal/domain"
)

type getByIdCategoryUseCase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type GetByIdCategoryUseCase interface {
	Perform(input GetByIdCategoryInput) (*GetByIdCategoryOutput, error)
}

func NewGetByIdCategoryUseCase(categoryRepo domain.CategoryRepository, userRepo domain.UserRepository) GetByIdCategoryUseCase {
	return &getByIdCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *getByIdCategoryUseCase) Perform(input GetByIdCategoryInput) (*GetByIdCategoryOutput, error) {
	if input.ID == "" {
		return nil, domain.ErrCategoryIdIsRequired
	}

	if input.UserId == "" {
		return nil, domain.ErrCategoryUserIdIsRequired
	}

	user, err := uc.userRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrCategoryUserNotFound
	}

	category, err := uc.categoryRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, domain.ErrCategoryNotFound
	}

	if category.UserId != user.ID {
		return nil, domain.ErrCategoryUserNotOwner
	}

	return &GetByIdCategoryOutput{
		ID:        category.ID,
		UserId:    category.UserId,
		Name:      category.Name,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
