package product

import (
	"github.com/areteacademy/internal/domain"
)

type listByUserIdProductUseCase struct {
	productRepo domain.ProductRepository
	userRepo    domain.UserRepository
}

type ListByUserIdProductUseCase interface {
	Perform(userId string) (ListByUserIdProductOutput, error)
}

func NewListByUserIdProductUseCase(
	productRepo domain.ProductRepository,
	userRepo domain.UserRepository,
) ListByUserIdProductUseCase {
	return &listByUserIdProductUseCase{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (uc *listByUserIdProductUseCase) Perform(userId string) (ListByUserIdProductOutput, error) {
	if userId == "" {
		return nil, domain.ErrProductUserIdIsRequired
	}

	user, err := uc.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrProductUserNotFound
	}

	producties, err := uc.productRepo.ListByUserId(userId)
	if err != nil {
		return nil, err
	}

	if producties == nil {
		return nil, domain.ErrProductNotFound
	}

	output := make(ListByUserIdProductOutput, 0, len(producties))

	for _, c := range producties {
		output = append(output, ProductItem{
			ID:          c.ID,
			UserId:      c.UserId,
			CategoryId:  c.CategoryId,
			Name:        c.Name,
			Description: c.Description,
			Status:      c.Status,
			Price:       c.Price,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}

	return output, nil
}
