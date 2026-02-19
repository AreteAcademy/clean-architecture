package user

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	repo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase UpdateUserUseCase
	Repo    *repo.InMemoryUserRepository
}

func makeSut() SUT {
	repo := repo.NewInMemoryUserRepository()
	usecase := NewUpdateUserUseCase(repo)

	return SUT{
		UseCase: usecase,
		Repo:    repo,
	}
}

func TestUpdateUser_ShouldReturnError_WhenIdIsEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domain.ErrUserIdIsRequired {
		t.Errorf("expected ErrUserIdIsRequired, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenNameIsEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "",
		Email: "daniel@gmail.com",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domain.ErrUserNameIsRequired {
		t.Errorf("expected ErrUserNameIsRequired, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenEmailIsEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Daniel",
		Email: "",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domain.ErrUserEmailIsRequired {
		t.Errorf("expected ErrUserEmailIsRequired, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenUserNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.Save(&domain.User{
		ID:    "1234",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domain.ErrUserNotFound {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenRepositoryFailOnGet(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.FailOnGet = true

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != repo.ErrSimulatedFailureRepoUser {
		t.Errorf("expected ErrSimulatedFailureRepoUser, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenEmailIsInvalid(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.Save(&domain.User{
		ID:    "123",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Daniel",
		Email: "danielgmail.com",
	})

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domain.ErrUserEmailInvalid {
		t.Errorf("expected ErrUserEmailInvalid, got %v", err)
	}
}

func TestUpdateUser_ShouldReturnError_WhenRepositoryFailOnUpdate(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.Repo.Save(&domain.User{
		ID:    "123",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})
	sut.Repo.FailOnUpdate = true

	// Act
	_, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != repo.ErrSimulatedFailureRepoUser {
		t.Errorf("expected ErrSimulatedFailureRepoUser, got %v", err)
	}
}

func TestUpdatUser_ShouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()
	now := time.Now()
	sut.Repo.Save(&domain.User{
		ID:        "123",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	})

	// Act
	user, err := sut.UseCase.Perform(UpdateUserInput{
		ID:    "123",
		Name:  "Updated",
		Email: "updated@gmail.com",
	})

	if err != nil {
		t.Fatalf("expected error, got nil")
	}

	if user.Name != "Updated" {
		t.Errorf("expected Updated, got %s", user.Name)
	}

	if user.Email != "updated@gmail.com" {
		t.Errorf("expected updated@gmail.com, got %s", user.Email)
	}

	if !user.CreatedAt.Equal(now) {
		t.Fatalf("expected CreatedAt to remain unchanged")
	}

	if !user.UpdatedAt.After(now) {
		t.Fatalf("expected UpdatedAt to be greater than original time")
	}
}
