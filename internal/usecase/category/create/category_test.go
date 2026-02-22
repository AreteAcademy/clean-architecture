package category

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      CreateCategoryUseCase
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewCreateCategoryUseCase(categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestCreateCategory_ShouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "",
		Name:   "Minha categoria",
		Status: "ACTIVE",
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

func TestCreateCategory_ShouldReturnAnError_WhenNameEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "",
		Status: "ACTIVE",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryNameIsRequired {
		t.Errorf("expected ErrCategoryNameIsRequired, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestCreateCategory_ShouldReturnAnError_WhenStatusEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryStatusIsRequired {
		t.Errorf("expected ErrCategoryStatusIsRequired, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestCreateCategory_ShouldReturnAnError_WhenStatusInvalid(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "INACTIV",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryStatusInvalid {
		t.Errorf("expected ErrCategoryStatusInvalid, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}

func TestCreateCategory_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
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
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "ACTIVE",
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

func TestCreateCategory_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
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

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "ACTIVE",
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

func TestCreateCategory_ShouldReturnAnError_WhenCategoryRepoFailOnSave(t *testing.T) {
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
	sut.CategoryRepo.FailOnSave = true

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "ACTIVE",
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

	count, err := sut.CategoryRepo.Count()
	if err != nil {
		t.Fatalf("unexpected error from Count: %v", err)
	}

	if count != 0 {
		t.Errorf("expected category to be saved, got %d", count)
	}
}

func TestCreateCategory_ShouldReturnSuccess(t *testing.T) {
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

	// Act
	category, err := sut.UseCase.Perform(CreateCategoryInput{
		UserId: "123456",
		Name:   "Categoria Daniel",
		Status: "ACTIVE",
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

	if !category.CreatedAt.Equal(category.UpdatedAt) {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal on creation")
	}

	count, err := sut.CategoryRepo.Count()
	if err != nil {
		t.Fatalf("unexpected error from Count: %v", err)
	}

	if count != 1 {
		t.Errorf("expected category to be saved, got %d", count)
	}
}
