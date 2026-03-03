package product

import (
	"testing"
	"time"

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
			category, err := sut.UseCase.Perform(tc.input)

			if err == nil && tc.expectedErr != nil {
				t.Errorf("expected error %v, got nil", tc.expectedErr)
			} else if err != nil && err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if category != nil {
				t.Errorf("expected nil category, got %+v", category)
			}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductCategoryNotFound {
		t.Errorf("expected ErrProductCategoryNotFound, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != categoryRepo.ErrSimulatedFailureRepoCategory {
		t.Errorf("expected ErrSimulatedFailureRepoCategory, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductCategoryUserNotOwner {
		t.Errorf("expected ErrProductCategoryUserNotOwner, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
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
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if product == nil {
		t.Fatalf("expected nil product, got %+v", product)
	}

	if product.ID == "" {
		t.Fatalf("expected ID to be set")
	}

	if product.UserId == "" {
		t.Fatalf("expected UserId to be set")
	}

	if product.UserId != "123456" {
		t.Fatalf("expected UserId, got %v", product.UserId)
	}

	if product.CategoryId == "" {
		t.Fatalf("expected CategoryId to be set")
	}

	if product.CategoryId != "123456" {
		t.Fatalf("expected CategoryId, got %v", product.CategoryId)
	}

	if product.Name != "Produto1" {
		t.Fatalf("expected Name Produto1, got %v", product.Name)
	}

	if product.Description != "Meu produto" {
		t.Fatalf("expected Description Meu produto, got %v", product.Description)
	}

	if product.Status != "ACTIVE" {
		t.Fatalf("expected Status ACTIVE, got %v", product.Status)
	}

	if product.Price != 100 {
		t.Fatalf("expected Price 100, got %v", product.Price)
	}

	if product.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be set")
	}

	if product.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt to be set")
	}

	if !product.CreatedAt.Equal(product.UpdatedAt) {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal on creation")
	}

	count, err := sut.ProductRepo.Count()
	if err != nil {
		t.Fatalf("unexpected error from Count: %v", err)
	}

	if count != 1 {
		t.Errorf("expected product to be saved, got %d", count)
	}
}
