package product

import (
	"github.com/areteacademy/internal/domain"
)

type updateProductUseCase struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type UpdateProductUseCase interface {
	Perform(input UpdateProductInput) (*UpdateProductOutput, error)
}

func NewUpdateProductUseCase(
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository,
	userRepo domain.UserRepository,
) UpdateProductUseCase {
	return &updateProductUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *updateProductUseCase) Perform(input UpdateProductInput) (*UpdateProductOutput, error) {
	product, err := uc.productRepo.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, domain.ErrProductNotFound
	}

	err = product.UpdateProduct(
		input.CategoryId,
		input.Name,
		input.Description,
		domain.ProductStatus(input.Status),
		input.Price,
	)

	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetById(product.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrProductUserNotFound
	}

	category, err := uc.categoryRepo.GetByIdAndUserId(input.CategoryId, product.UserId)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, domain.ErrProductCategoryNotFound
	}

	if err := uc.productRepo.Update(product); err != nil {
		return nil, err
	}

	return &UpdateProductOutput{
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
