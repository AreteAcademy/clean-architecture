package category

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      UpdateCategoryUseCase
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewUpdateCategoryUseCase(categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestUpdateCategory_ShouldReturnAnError_WhenInputValidators(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	testCases := []struct {
		name        string
		input       UpdateCategoryInput
		expectedErr error
	}{
		{
			name: "Empty ID",
			input: UpdateCategoryInput{
				ID:     "",
				UserId: "123456",
				Name:   "Category1",
				Status: "ACTIVE",
			},
			expectedErr: domain.ErrCategoryIdIsRequired,
		},
		{
			name: "Empty User ID",
			input: UpdateCategoryInput{
				ID:     "123456",
				UserId: "",
				Name:   "Category1",
				Status: "ACTIVE",
			},
			expectedErr: domain.ErrCategoryUserIdIsRequired,
		},
		{
			name: "Empty Name",
			input: UpdateCategoryInput{
				ID:     "123456",
				UserId: "123456",
				Name:   "",
				Status: "ACTIVE",
			},
			expectedErr: domain.ErrCategoryNameIsRequired,
		},
		{
			name: "Empty Status",
			input: UpdateCategoryInput{
				ID:     "123456",
				UserId: "123456",
				Name:   "Category1",
				Status: "",
			},
			expectedErr: domain.ErrCategoryStatusIsRequired,
		},
		{
			name: "Invalid Status",
			input: UpdateCategoryInput{
				ID:     "123456",
				UserId: "123456",
				Name:   "Category1",
				Status: "ACTIV",
			},
			expectedErr: domain.ErrCategoryStatusInvalid,
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

func TestUpdateCategory_ShouldReturnAnError_WhenCategoryNotFound(t *testing.T) {
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
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
		Status: "ACTIVE",
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

func TestUpdateCategory_ShouldReturnAnError_WhenCategoryRepoFailOnGetById(t *testing.T) {
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
	sut.CategoryRepo.FailOnGetById = true

	// Act
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
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
}

func TestUpdateCategory_ShouldReturnAnError_WhenCategoryUserIdNotOwner(t *testing.T) {
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
		UserId:    "1234567",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
		Status: "ACTIVE",
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

func TestUpdateCategory_ShouldReturnAnError_WhenCategoryUserNotFound(t *testing.T) {
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
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
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

func TestUpdateCategory_ShouldReturnAnError_WhenUserRepoFailOnGetById(t *testing.T) {
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
		UserId:    "123456",
		Name:      "Categoria1",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
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

func TestUpdateCategory_ShouldReturnAnError_WhenCategoryRepoFailOnUpdate(t *testing.T) {
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
	sut.CategoryRepo.FailOnUpdate = true

	// Act
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
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
}

func TestUpdateCategory_ShouldReturnSuccess(t *testing.T) {
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
		Name:      "Categoria",
		Status:    "ACTIVE",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	category, err := sut.UseCase.Perform(UpdateCategoryInput{
		ID:     "123456",
		UserId: "123456",
		Name:   "Categoria editada",
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

	if category.Name != "Categoria editada" {
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

	if !category.UpdatedAt.After(now) {
		t.Fatalf("expected UpdatedAt to be greater than original time")
	}
}
