package user

import (
	"testing"

	"github.com/areteacademy/internal/domain"
)

// INFRA
type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetById(id string) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// SYSTEM UNDER TEST
type SUT struct {
	UseCase CreateUserUseCase
	Repo    *InMemoryUserRepository
}

func makeSut() SUT {
	repo := NewInMemoryUserRepository()
	usecase := NewCreateUserUseCase(repo)

	return SUT{
		UseCase: usecase,
		Repo:    repo,
	}
}

// TEST

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

func TestCreateUser_ShouldReturnAnErrorIfEmailEmpty(t *testing.T) {
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

func TestCreateUser_ShouldReturnAnErrorIfEmailInvalid(t *testing.T) {
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

func TestCreateUser_ShouldReturnAnErrorIfPasswordEmpty(t *testing.T) {
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

func TestCreateUser_ShouldReturnAnErrorIfPasswordInvalid(t *testing.T) {
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
}

func TestCreateUser_ShouldSaveOnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, err := sut.UseCase.Perform(&CreateUserInput{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	if err != nil {
		t.Fatalf("expected an error, got nil")
	}

	if len(sut.Repo.users) != 1 {
		t.Errorf("expected user to be saved")
	}
}

func TestCreateUser_shouldNotSaveWhenValidationFails(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, _ = sut.UseCase.Perform(&CreateUserInput{
		Name:     "",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	if len(sut.Repo.users) != 0 {
		t.Errorf("expected not user to be saved")
	}

}
