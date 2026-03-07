package product

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	productRepo "github.com/areteacademy/internal/infra/repository/product"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      CreateProductUseCase
	ProductRepo  *productRepo.InMemoryProductRepository
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	productRepo := productRepo.NewInMemoryProductRepository()
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewCreateProductUseCase(productRepo, categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		ProductRepo:  productRepo,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestCreateProduct_ShouldReturnAnError_WhenInputValidators(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	testCases := []struct {
		name        string
		input       CreateProductInput
		expectedErr error
	}{
		{
			name: "Empty User ID",
			input: CreateProductInput{
				UserId:      "",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "ACTIVE",
				Price:       100,
			},
			expectedErr: domain.ErrProductUserIdIsRequired,
		},
		{
			name: "Empty Category ID",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "ACTIVE",
				Price:       100,
			},
			expectedErr: domain.ErrProductCategoryIdIsRequired,
		},
		{
			name: "Empty Name",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "",
				Description: "Meu produto",
				Status:      "ACTIVE",
				Price:       100,
			},
			expectedErr: domain.ErrProductNameIsRequired,
		},
		{
			name: "Empty Description",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "",
				Status:      "ACTIVE",
				Price:       100,
			},
			expectedErr: domain.ErrProductDescriptionIsRequired,
		},
		{
			name: "Empty Status",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "",
				Price:       100,
			},
			expectedErr: domain.ErrProductStatusIsRequired,
		},
		{
			name: "Invalid Status",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "INACTI",
				Price:       100,
			},
			expectedErr: domain.ErrProductStatusInvalid,
		},
		{
			name: "Invalid Price",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "ACTIVE",
				Price:       0,
			},
			expectedErr: domain.ErrProductPriceInvalid,
		},
		{
			name: "Invalid Price",
			input: CreateProductInput{
				UserId:      "123456",
				CategoryId:  "123456",
				Name:        "Produto1",
				Description: "Meu produto",
				Status:      "ACTIVE",
				Price:       -100,
			},
			expectedErr: domain.ErrProductPriceInvalid,
		},
	}

	// Assert
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			product, err := sut.UseCase.Perform(tc.input)

			require.Error(t, err)
			require.Nil(t, product)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestCreateProduct_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "1234567",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductUserNotFound)
}

func TestCreateProduct_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "1234567",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.UserRepo.FailOnGet = true

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, userRepo.ErrSimulatedFailureRepoUser)
}

func TestCreateProduct_ShouldReturnAnError_WhenCategoryNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "1234567",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductCategoryNotFound)
}

func TestCreateProduct_ShouldReturnAnError_WhenCategoryRepoFailOnGetById(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.CategoryRepo.FailOnGet = true

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, categoryRepo.ErrSimulatedFailureRepoCategory)
}

func TestCreateProduct_ShouldReturnAnError_WhenCategoryUserNotOwner(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.UserRepo.Save(&domain.User{
		ID:        "1234567",
		Name:      "Daniel2",
		Email:     "daniel2@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.CategoryRepo.Save(&domain.Category{
		ID:        "1234567",
		UserId:    "1234567",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "1234567",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductCategoryUserNotOwner)
}

func TestCreateProduct_ShouldReturnAnError_WhenProductRepoFailOnSave(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.ProductRepo.FailOnSave = true

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, productRepo.ErrSimulatedFailureRepoProduct)
}

func TestCreateProduct_ShouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.UserRepo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	require.NoError(t, err)
	require.NotNil(t, product)

	assert.NotNil(t, product.ID)
	assert.Equal(t, "123456", product.UserId)
	assert.Equal(t, "123456", product.CategoryId)
	assert.Equal(t, "Produto1", product.Name)
	assert.Equal(t, "Meu Produto", product.Description)
	assert.Equal(t, "ACTIVE", product.Status)
	assert.Equal(t, 100, product.Price)

	assert.False(t, product.CreatedAt.IsZero())
	assert.False(t, product.UpdatedAt.IsZero())

	count, err := sut.ProductRepo.Count()
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}
