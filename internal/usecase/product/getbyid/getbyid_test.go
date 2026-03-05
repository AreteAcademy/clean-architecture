package product

import (
	"testing"
	"time"

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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductIdIsRequired {
		t.Errorf("expected ErrProductIdIsRequired, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductUserIdIsRequired {
		t.Errorf("expected ErrProductUserIdIsRequired, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductUserNotFound {
		t.Errorf("expected ErrProductUserNotFound, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != userRepo.ErrSimulatedFailureRepoUser {
		t.Errorf("expected ErrSimulatedFailureRepoUser, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductNotFound {
		t.Errorf("expected ErrProductNotFound, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
}

func TestGetByIdProduct_ShouldReturnAnError_WhenProductRepoFailOnGetById(t *testing.T) {
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
	sut.ProductRepo.FailOnGetById = true

	// Act
	product, err := sut.UseCase.Perform(GetByIdProductInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != productRepo.ErrSimulatedFailureRepoProduct {
		t.Errorf("expected ErrSimulatedFailureRepoProduct, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err != nil {
		t.Fatalf("unexptected error: %v", err)
	}

	if product == nil {
		t.Fatalf("expected nil product, got %+v", product)
	}

	if product.ID == "" {
		t.Errorf("expected ID to be set")
	}

	if product.UserId == "" {
		t.Errorf("expected UserId to be set")
	}

	if product.CategoryId == "" {
		t.Errorf("expected CategoryId to be set")
	}

	if product.CategoryId != "123456" {
		t.Errorf("expected CategoryId to be set value 123456")
	}

	if product.Name == "" {
		t.Errorf("expected Name to be set")
	}

	if product.Description == "" {
		t.Errorf("expected Description to be set")
	}

	if product.Status == "" {
		t.Errorf("expected Description to be set")
	}

	if product.Price == 0 {
		t.Errorf("expected Price to be set")
	}

	if product.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be set")
	}

	if product.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt to be set")
	}
}
