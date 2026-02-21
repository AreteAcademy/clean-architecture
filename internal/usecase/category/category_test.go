package category

import (
	"errors"
	"testing"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

var (
	ErrCategoryUserIdIsRequired = errors.New("user id is required")
)

type CreateCategoryInput struct {
	UserId string
	Name   string
	Status string
}

type CreateCategoryOutput struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

type createCategoryUseCase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

type CreateCategoryUseCase interface {
	Perform(input CreateCategoryInput) (*CreateCategoryOutput, error)
}

func NewCreateCategoryUseCase(categoryRepo domain.CategoryRepository, userRepo domain.UserRepository) CreateCategoryUseCase {
	return &createCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc *createCategoryUseCase) Perform(input CreateCategoryInput) (*CreateCategoryOutput, error) {
	if input.UserId == "" {
		return nil, ErrCategoryUserIdIsRequired
	}

	return &CreateCategoryOutput{}, nil
}

type SUT struct {
	UseCase      CreateCategoryUseCase
	CategoryRepo domain.CategoryRepository
	UserRepo     domain.UserRepository
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

func TestCreateCategory_ShouldReturnAnError_WhenIdUserIdEmpty(t *testing.T) {
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

	if err != ErrCategoryUserIdIsRequired {
		t.Errorf("expected ErrCategoryUserIdIsRequired, got %v", err)
	}

	if category != nil {
		t.Errorf("expected nil category, got %+v", category)
	}
}
