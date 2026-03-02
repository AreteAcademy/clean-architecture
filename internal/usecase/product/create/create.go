package product

import "github.com/areteacademy/internal/domain"

type createProductUseCase struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type CreateProductUseCase interface {
	Perform(input CreateProductInput) (*CreateProductOutput, error)
}

func NewCreateProductUseCase(
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository,
	userRepo domain.UserRepository,
) CreateProductUseCase {
	return &createProductUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *createProductUseCase) Perform(input CreateProductInput) (*CreateProductOutput, error) {
	if input.UserId == "" {
		return nil, domain.ErrProductUserIdIsRequired
	}

	return nil, nil
}
