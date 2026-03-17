package product

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	productRepo "github.com/areteacademy/internal/infra/repository/product"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SUT struct {
	UseCase      UpdateProductUseCase
	ProductRepo  *productRepo.InMemoryProductRepository
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
	User         *domain.User
	Category     *domain.Category
	Product      *domain.Product
}

func makeSut() SUT {
	productRepo := productRepo.NewInMemoryProductRepository()
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewUpdateProductUseCase(productRepo, categoryRepo, userRepo)

	now := time.Now()
	user := &domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	}
	category := &domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	}
	product := &domain.Product{
		ID:          "123456",
		UserId:      user.ID,
		CategoryId:  category.ID,
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return SUT{
		UseCase:      usecase,
		ProductRepo:  productRepo,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
		User:         user,
		Category:     category,
		Product:      product,
	}
}

func validInput(sut SUT) UpdateProductInput {
	return UpdateProductInput{
		ID:          sut.Product.ID,
		UserId:      sut.User.ID,
		CategoryId:  sut.Product.CategoryId,
		Name:        sut.Product.Name,
		Description: sut.Product.Description,
		Status:      sut.Product.Status,
		Price:       sut.Product.Price,
	}
}

func seedDefaultData(sut SUT) {
	sut.UserRepo.Save(sut.User)
	sut.CategoryRepo.Save(sut.Category)
	sut.ProductRepo.Save(sut.Product)
}

func TestUpdateProduct_GivenInvalidInput_ShouldReturnError(t *testing.T) {
	// Act
	testCases := []struct {
		name        string
		setup       func(sut SUT)
		input       func(sut SUT) UpdateProductInput
		expectedErr error
	}{
		{
			name: "Empty ID",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.ID = ""
				return in
			},
			expectedErr: domain.ErrProductIdIsRequired,
		},
		{
			name: "Empty User ID",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.UserId = ""
				return in
			},
			expectedErr: domain.ErrProductUserIdIsRequired,
		},
		{
			name: "Empty Category ID",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.CategoryId = ""
				return in
			},
			expectedErr: domain.ErrProductCategoryIdIsRequired,
		},
		{
			name: "Empty Name",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Name = ""
				return in
			},
			expectedErr: domain.ErrProductNameIsRequired,
		},
		{
			name: "Empty Description",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Description = ""
				return in
			},
			expectedErr: domain.ErrProductDescriptionIsRequired,
		},
		{
			name: "Empty Status",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Status = ""
				return in
			},
			expectedErr: domain.ErrProductStatusIsRequired,
		},
		{
			name: "Invalid Status",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Status = "ACTIV"
				return in
			},
			expectedErr: domain.ErrProductStatusInvalid,
		},
		{
			name: "Invalid Price",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Price = -10
				return in
			},
			expectedErr: domain.ErrProductPriceInvalid,
		},
		{
			name: "Invalid Price",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.Price = 0
				return in
			},
			expectedErr: domain.ErrProductPriceInvalid,
		},
		{
			name: "Product Not Found",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.ID = "1234567"
				return in
			},
			expectedErr: domain.ErrProductNotFound,
		},
		{
			name: "Repo Product Fail On GetByIdAndUserId",
			setup: func(sut SUT) {
				seedDefaultData(sut)
				sut.ProductRepo.FailOnGetByIdAndUserId = true
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				return in
			},
			expectedErr: productRepo.ErrSimulatedFailureRepoProduct,
		},
		{
			name: "Repo User Fail On GetById",
			setup: func(sut SUT) {
				seedDefaultData(sut)
				sut.UserRepo.FailOnGet = true
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				return in
			},
			expectedErr: userRepo.ErrSimulatedFailureRepoUser,
		},
		{
			name: "Category Not Found",
			setup: func(sut SUT) {
				seedDefaultData(sut)
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				in.CategoryId = "1234567"
				return in
			},
			expectedErr: domain.ErrProductCategoryNotFound,
		},
		{
			name: "Repo Category Fail On GetByIdAndUserId",
			setup: func(sut SUT) {
				seedDefaultData(sut)
				sut.CategoryRepo.FailOnGetByIdAndUserId = true
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				return in
			},
			expectedErr: categoryRepo.ErrSimulatedFailureRepoCategory,
		},
		{
			name: "Repo Product Fail On Update",
			setup: func(sut SUT) {
				seedDefaultData(sut)
				sut.ProductRepo.FailOnUpdate = true
			},
			input: func(sut SUT) UpdateProductInput {
				in := validInput(sut)
				return in
			},
			expectedErr: productRepo.ErrSimulatedFailureRepoProduct,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			sut := makeSut()

			if tc.setup != nil {
				tc.setup(sut)
			}

			input := tc.input(sut)

			product, err := sut.UseCase.Perform(input)

			// Assert
			require.Error(t, err)
			require.Nil(t, product)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestUpdateProduct_ShouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()
	seedDefaultData(sut)

	expected := UpdateProductInput{
		ID:          sut.Product.ID,
		UserId:      sut.User.ID,
		CategoryId:  sut.Product.CategoryId,
		Name:        "Produto editado",
		Description: "Descrição editado",
		Status:      sut.Product.Status,
		Price:       190,
	}

	// Act
	product, err := sut.UseCase.Perform(expected)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, product)

	assert.Equal(t, expected.ID, product.ID)
	assert.Equal(t, sut.Product.UserId, product.UserId)
	assert.Equal(t, expected.CategoryId, product.CategoryId)
	assert.Equal(t, expected.Name, product.Name)
	assert.Equal(t, expected.Description, product.Description)
	assert.Equal(t, expected.Status, product.Status)
	assert.Equal(t, expected.Price, product.Price)
	assert.False(t, product.UpdatedAt.IsZero())
	assert.True(t, product.UpdatedAt.After(product.CreatedAt))
}
