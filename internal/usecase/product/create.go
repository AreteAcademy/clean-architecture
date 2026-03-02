package product

import "github.com/areteacademy/internal/domain"

type createProductUseCase struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type CreateCategoryUseCase interface {
	Perform(input CreateProductInput) (*CreateProductOutput, error)
}

func NewCreateProductUseCase(
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository,
	userRepo domain.UserRepository,
) CreateCategoryUseCase {
	return &createProductUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *createProductUseCase) Perform(input CreateProductInput) (*CreateProductOutput, error) {
	if input.CategoryId == "" {
		return nil, domain.ErrProductCategoryIdIsRequired
	}

	return nil, nil
}
