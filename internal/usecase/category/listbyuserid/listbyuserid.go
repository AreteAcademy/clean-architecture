package category

import (
	"github.com/areteacademy/internal/domain"
)

type listByUserIdCategoryUseCase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type ListByUserIdCategoryUseCase interface {
	Perform(userId string) (ListUserByIdCategoryOutput, error)
}

func NewListByUserIdCategoryUseCase(
	categoryRepo domain.CategoryRepository,
	userRepo domain.UserRepository,
) ListByUserIdCategoryUseCase {
	return &listByUserIdCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *listByUserIdCategoryUseCase) Perform(userId string) (ListUserByIdCategoryOutput, error) {
	if userId == "" {
		return nil, domain.ErrCategoryUserIdIsRequired
	}

	user, err := uc.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrCategoryUserNotFound
	}

	categories, err := uc.categoryRepo.ListByUserId(userId)
	if err != nil {
		return nil, err
	}

	output := make(ListUserByIdCategoryOutput, 0, len(categories))

	for _, c := range categories {
		output = append(output, CategoryItem{
			ID:        c.ID,
			UserId:    c.UserId,
			Name:      c.Name,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return output, nil
}
