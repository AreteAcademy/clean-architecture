package user

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	repo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase GetByIdUserUseCase
	Repo    *repo.InMemoryUserRepository
}

func makeSut() SUT {
	repo := repo.NewInMemoryUserRepository()
	usecase := NewGetByIdUserUseCase(repo)

	return SUT{
		UseCase: usecase,
		Repo:    repo,
	}
}

func TestGetById_shouldReturnAnErrorIfNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform("123456")

	// Assert
	if err == nil {
		t.Fatal("expected an error, not nil")
	}

	if err != domain.ErrUserNotFound {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func TestGetById_shouldReturnUserSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.Repo.Save(&domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	user, err := sut.UseCase.Perform("123456")

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Name != "Daniel" {
		t.Errorf("expected name Daniel, got %s", user.Name)
	}

	if user.Email != "daniel@gmail.com" {
		t.Errorf("expected name daniel@gmail.com, got %s", user.Email)
	}

	if user.ID != "123456" {
		t.Errorf("expected id 123456, got %s", user.ID)
	}

	if user.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set")
	}

	if user.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set")
	}
}

func TestGetById_shouldReturnAnErrorIfIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform("")

	// Assert
	if err == nil {
		t.Fatal("expected an error, not nil")
	}

	if err != domain.ErrUserIdIsRequired {
		t.Errorf("expected ErrUserIdIsRequired, got %v", err)
	}
}

func TestGetById_shouldReturnErrorWhenRepositoryFail(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.Save(&domain.User{
		ID:    "123456",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})
	sut.Repo.FailOnGet = true

	// Act
	user, err := sut.UseCase.Perform("123456")

	// Assert
	if err == nil {
		t.Fatalf("expected an error, not nil")
	}

	if err != repo.ErrSimulatedFailureRepoUser {
		t.Errorf("expected repository erro, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}
