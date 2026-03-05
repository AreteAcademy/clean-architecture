package product

import (
	"github.com/areteacademy/internal/domain"
)

type getByIdProductUseCase struct {
	productRepo domain.ProductRepository
	userRepo    domain.UserRepository
}

type GetByIdProductUseCase interface {
	Perform(input GetByIdProductInput) (*GetByIdProductOutput, error)
}

func NewGetByIdProductUseCase(
	productRepo domain.ProductRepository,
	userRepo domain.UserRepository,
) GetByIdProductUseCase {
	return &getByIdProductUseCase{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (uc *getByIdProductUseCase) Perform(input GetByIdProductInput) (*GetByIdProductOutput, error) {
	if input.ID == "" {
		return nil, domain.ErrProductIdIsRequired
	}

	if input.UserId == "" {
		return nil, domain.ErrProductUserIdIsRequired
	}

	user, err := uc.userRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrProductUserNotFound
	}

	product, err := uc.productRepo.GetByIdAndUserId(input.ID, input.UserId)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, domain.ErrProductNotFound
	}

	return &GetByIdProductOutput{
		ID:          product.ID,
		UserId:      product.UserId,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Description: product.Description,
		Status:      product.Status,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
