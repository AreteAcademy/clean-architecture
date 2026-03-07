package product

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/areteacademy/internal/domain"
	productRepo "github.com/areteacademy/internal/infra/repository/product"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase     GetByIdProductUseCase
	ProductRepo *productRepo.InMemoryProductRepository
	UserRepo    *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	productRepo := productRepo.NewInMemoryProductRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewGetByIdProductUseCase(productRepo, userRepo)

	return SUT{
		UseCase:     usecase,
		ProductRepo: productRepo,
		UserRepo:    userRepo,
	}
}

func TestGetByIdProduct_ShouldReturnAnError_WhenIdEmpty(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "",
		UserId: "123456",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductIdIsRequired)
}

func TestGetByIdProduct_ShouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123455",
		UserId: "",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductUserIdIsRequired)
}

func TestGetByIdProduct_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductUserNotFound)
}

func TestGetByIdProduct_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, userRepo.ErrSimulatedFailureRepoUser)
}

func TestGetByIdProduct_ShouldReturnAnError_WhenProductNotFound(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "1234567",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrProductNotFound)
}

func TestGetByIdProduct_ShouldReturnAnError_WhenProductRepoFailOnGetByIdAnduserId(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	sut.ProductRepo.FailOnGetByIdAndUserId = true

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	require.Error(t, err)
	require.Nil(t, product)
	assert.ErrorIs(t, err, productRepo.ErrSimulatedFailureRepoProduct)
}

func TestGetByIdProduct_ShouldReturnSuccess(t *testing.T) {
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

	sut.ProductRepo.Save(&domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	require.NoError(t, err)
	require.NotNil(t, product)

	assert.Equal(t, "123456", product.ID)
	assert.Equal(t, "123456", product.UserId)
	assert.Equal(t, "123456", product.CategoryId)
	assert.Equal(t, "Produto1", product.Name)
	assert.Equal(t, "Meu Produto", product.Description)
	assert.Equal(t, "ACTIVE", product.Status)
	assert.Equal(t, 100, product.Price)
	assert.Equal(t, now, product.CreatedAt)
	assert.Equal(t, now, product.UpdatedAt)

	assert.False(t, product.CreatedAt.IsZero())
	assert.False(t, product.UpdatedAt.IsZero())
}
