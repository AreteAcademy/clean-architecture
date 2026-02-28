package category

import "github.com/areteacademy/internal/domain"

type updateCategoryUseCase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type UpdateCategoryUseCase interface {
	Perform(input UpdateCategoryInput) (*UpdateCategoryOutput, error)
}

func NewUpdateCategoryUseCase(categoryRepo domain.CategoryRepository, userRepo domain.UserRepository) UpdateCategoryUseCase {
	return &updateCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *updateCategoryUseCase) Perform(input UpdateCategoryInput) (*UpdateCategoryOutput, error) {
	category, err := domain.UpdateCategory(
		input.ID,
		input.UserId,
		input.Name,
		domain.CategoryStatus(input.Status),
	)

	if err != nil {
		return nil, err
	}

	exists, err := uc.categoryRepo.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, domain.ErrCategoryNotFound
	}

	if exists.UserId != input.UserId {
		return nil, domain.ErrCategoryUserNotOwner
	}

	user, err := uc.userRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrCategoryUserNotFound
	}

	if err := uc.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return &UpdateCategoryOutput{
		ID:        category.ID,
		UserId:    category.UserId,
		Name:      category.Name,
		Status:    category.Status,
		CreatedAt: exists.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
