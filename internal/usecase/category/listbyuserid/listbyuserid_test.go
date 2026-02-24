package category

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      ListByUserIdCategoryUseCase
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewListByUserIdCategoryUseCase(categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestListByUserIdCategory_shouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
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
	categories, err := sut.UseCase.Perform("")

	// Assert
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if categories != nil {
		t.Errorf("expected nil categories, got %+v", categories)
	}

	if err != domain.ErrCategoryUserIdIsRequired {
		t.Errorf("expected ErrCategoryUserIdIsRequired, got %v", err)
	}
}

func TestListByUserIdCategory_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
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
		Name:      "Categoria Daniel",
		UserId:    "123456",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	categories, err := sut.UseCase.Perform("123456")

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrCategoryUserNotFound {
		t.Errorf("expected ErrCategoryUserNotFound, got %v", err)
	}

	if categories != nil {
		t.Errorf("expected nil category, got %+v", categories)
	}
}

func TestLisByUserIdCategory_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
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

	sut.CategoryRepo.Save(&domain.Category{
		ID:        "123456",
		Name:      "Categoria Daniel",
		UserId:    "123456",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	categories, err := sut.UseCase.Perform("123456")

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != userRepo.ErrSimulatedFailureRepoUser {
		t.Errorf("expected ErrSimulatedFailureRepoUser, got %v", err)
	}

	if categories != nil {
		t.Errorf("expected nil category, got %+v", categories)
	}
}

func TestLisByUserIdCategory_ShouldReturnAnError_WhenCategoryRepoFailOnListByUserId(t *testing.T) {
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
	sut.CategoryRepo.FailOnList = true

	// Act
	categories, err := sut.UseCase.Perform("123456")

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != categoryRepo.ErrSimulatedFailureRepoCategory {
		t.Errorf("expected ErrSimulatedFailureRepoCategory, got %v", err)
	}

	if categories != nil {
		t.Errorf("expected nil categories, got %+v", categories)
	}
}

func TestLisByUserIdCategory_ShouldReturnSuccess(t *testing.T) {
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
		ID:        "cat1",
		Name:      "cat1",
		UserId:    "123456",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.CategoryRepo.Save(&domain.Category{
		ID:        "cat2",
		Name:      "cat2",
		UserId:    "123456",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})
	sut.CategoryRepo.Save(&domain.Category{
		ID:        "cat3",
		Name:      "cat3",
		UserId:    "999999",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	categories, err := sut.UseCase.Perform("123456")

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if categories == nil {
		t.Fatalf("expected categories, got nil")
	}

	if len(categories) != 2 {
		t.Errorf("expected two categories, got %d", len(categories))
	}

	for _, c := range categories {
		if c.UserId != "123456" {
			t.Errorf("expected category to belong to user 123456, got %s", c.UserId)
		}

		if c.ID == "" {
			t.Errorf("expected ID to be set")
		}

		if c.Status != "ACTIVE" {
			t.Errorf("expected status ACTIVE, got %s", c.Status)
		}

		if c.CreatedAt.IsZero() {
			t.Errorf("expected CreatedAt to be set")
		}

		if c.UpdatedAt.IsZero() {
			t.Errorf("expected UpdatedAt to be set")
		}
	}
}
