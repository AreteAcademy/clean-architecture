package user

import (
	"testing"

	"github.com/areteacademy/internal/domain"
	repo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase CreateUserUseCase
	Repo    *repo.InMemoryUserRepository
}

func makeSut() SUT {
	repo := repo.NewInMemoryUserRepository()
	usecase := NewCreateUserUseCase(repo)

	return SUT{
		UseCase: usecase,
		Repo:    repo,
	}
}

func TestCreateUser_ShouldReturnAnErrorIfNameEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "",
		Email:    "daniel@gmail.com",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrUserNameIsRequired {
		t.Errorf("expected ErrUserNameIsRequired, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}

func TestCreateUser_ShouldReturnAnError_WhenIfEmailEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrUserEmailIsRequired {
		t.Errorf("expected ErrUserEmailIsRequired, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}

func TestCreateUser_ShouldReturnAnError_WhenIfEmailInvalid(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel.com.br",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrUserEmailInvalid {
		t.Errorf("expected ErrUserEmailInvalid, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}

func TestCreateUser_ShouldReturnAnError_WhenIfPasswordEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrUserPasswordIsRequired {
		t.Errorf("expected ErrUserPasswordIsRequired, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}

func TestCreateUser_ShouldReturnAnError_WhenIfPasswordInvalid(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrUserPasswordInvalid {
		t.Errorf("expected ErrUserPasswordInvalid, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}

func TestCreateUser_ShouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	if err != nil {
		t.Fatalf("expected an error, got nil")
	}

	if user == nil {
		t.Errorf("expected nil user, got %+v", user)
	}

	if user.Name != "Daniel" || user.Email != "daniel@gmail.com" {
		t.Errorf("expected nil user return comform, got %+v", user)
	}

	if user.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set")
	}

	if user.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set")
	}

	if !user.CreatedAt.Equal(user.UpdatedAt) {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal on creation")
	}
}

func TestCreateUser_ShouldSaveOnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, _ = sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	count, err := sut.Repo.Count()
	if err != nil {
		t.Fatalf("unexpected error from Count: %v", err)
	}

	if count != 1 {
		t.Errorf("expected user to be saved, got %d", count)
	}
}

func TestCreateUser_shouldReturnAnError_WhenValidationFails(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, _ = sut.UseCase.Perform(&CreateUserInput{
		Name:     "",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	count, err := sut.Repo.Count()
	if err != nil {
		t.Fatalf("unexpected error from Count: %v", err)
	}

	if count != 0 {
		t.Errorf("expected user to be saved, got %d", count)
	}
}

func TestCreateUser_shouldReturnAnError_WhenRepositoryFails(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.FailOnSave = true

	// Act
	user, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

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
