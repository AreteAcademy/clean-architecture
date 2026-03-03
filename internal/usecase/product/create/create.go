package product

import (
	"github.com/areteacademy/internal/domain"
)

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
	product, err := domain.NewProduct(
		input.UserId,
		input.CategoryId,
		input.Name,
		input.Description,
		domain.ProductStatus(input.Status),
		input.Price,
	)

	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetById(input.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrProductUserNotFound
	}

	category, err := uc.categoryRepo.GetById(input.CategoryId)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, domain.ErrProductCategoryNotFound
	}

	if category.UserId != input.UserId {
		return nil, domain.ErrProductCategoryUserNotOwner
	}

	if err := uc.productRepo.Save(product); err != nil {
		return nil, err
	}

	return &CreateProductOutput{
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
