package category

import "github.com/areteacademy/internal/domain"

type createCategoryUseCase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type CreateCategoryUseCase interface {
	Perform(input CreateCategoryInput) (*CreateCategoryOutput, error)
}

func NewCreateCategoryUseCase(categoryRepo domain.CategoryRepository, userRepo domain.UserRepository) CreateCategoryUseCase {
	return &createCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *createCategoryUseCase) Perform(input CreateCategoryInput) (*CreateCategoryOutput, error) {
	category, err := domain.NewCategory(
		input.UserId,
		input.Name,
		domain.CategoryStatus(input.Status),
	)

	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrCategoryUserNotFound
	}

	if err := uc.categoryRepo.Save(category); err != nil {
		return nil, err
	}

	return &CreateCategoryOutput{
		ID:        category.ID,
		UserId:    category.UserId,
		Name:      category.Name,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
