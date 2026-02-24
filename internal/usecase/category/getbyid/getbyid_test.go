package category

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      GetByIdCategoryUseCase
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewGetByIdCategoryUseCase(categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryIdIsRequired {
		t.Errorf("expected ErrCategoryIdIsRequired, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryUserIdIsRequired {
		t.Errorf("expected ErrCategoryUserIdIsRequired, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
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

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryUserNotFound {
		t.Errorf("expected ErrCategoryUserNotFound, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenCategoryNotFound(t *testing.T) {
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
		Name:      "Categoria Daniel",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryNotFound {
		t.Errorf("expected ErrCategoryNotFound, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
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
	sut.UserRepo.FailOnGet = true

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
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

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenCategoryRepoFailOnGetById(t *testing.T) {
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
		Name:      "Categoria Daniel",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.CategoryRepo.FailOnGet = true

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != categoryRepo.ErrSimulatedFailureRepoCategory {
		t.Errorf("expected ErrSimulatedFailureRepoCategory, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnAnError_WhenCategoryNotOwner(t *testing.T) {
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
		Name:      "Daniel 2",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		UserId:    "1234567",
		Name:      "Categoria Daniel",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryUserNotOwner {
		t.Errorf("expected ErrCategoryUserNotOwner, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestGetByIdCategory_ShouldReturnSuccess(t *testing.T) {
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
		Name:      "Categoria Daniel",
		UserId:    "123456",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(GetByIdCategoryInput{
		ID:     "123456",
		UserId: "123456",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if category == nil {
		t.Errorf("expected nil category, got %+v", category)
	}

	if category.ID == "" {
		t.Fatalf("expected ID to be set")
	}

	if category.UserId == "" {
		t.Fatalf("expected UserId to be set")
	}

	if category.Name != "Categoria Daniel" {
		t.Errorf("expected Category Name, got %v", category.Name)
	}

	if category.Status != "ACTIVE" {
		t.Errorf("expected Category Status, got %v", category.Status)
	}

	if category.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be set")
	}

	if category.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt to be set")
	}
}
