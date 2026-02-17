package getbyid

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

type getByIdUserUseCase struct {
	repo domain.UserRepository
}

type GetByIdUserUseCase interface {
	Perform(id string) (*GetByIdOutput, error)
}

func NewGetByIdUserUseCase(repo domain.UserRepository) GetByIdUserUseCase {
	return &getByIdUserUseCase{
		repo: repo,
	}
}

func (uc *getByIdUserUseCase) Perform(id string) (*GetByIdOutput, error) {
	user, err := uc.repo.GetById(id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &GetByIdOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// SYSTEM UNDER TEST
type SUT struct {
	UseCase GetByIdUserUseCase
	Repo    *InMemoryUserRepository
}

func makeSut() SUT {
	repo := NewInMemoryUserRepository()
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
	sut.Repo.users["123456"] = &domain.User{
		ID:    "123456",
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	}

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
}
