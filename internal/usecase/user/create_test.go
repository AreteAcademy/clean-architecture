package user_test

import (
	"errors"
	"testing"
)

// INTERFACES

var (
	ErrUserNameIsRequired = errors.New("name is required")
)

type User struct {
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user *User) error
}

type CreateUserUseCase struct {
	repo UserRepository
}

// CONSTRUCTOR

func NewCreateUserUseCase(repo UserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		repo: repo,
	}
}

func (uc *CreateUserUseCase) Perform(user *User) (*User, error) {
	if user.Name == "" {
		return nil, ErrUserNameIsRequired
	}

	err := uc.repo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// INFRA
type InMemoryUserRepository struct {
	users []*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []*User{},
	}
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users = append(r.users, user)
	return nil
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

func TestCreateUser_ShouldReturnAnErrorIsNameEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&User{
		Name:     "",
		Email:    "daniel@gmail.com",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != ErrUserNameIsRequired {
		t.Errorf("expected ErrUserNameIsRequired, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
}
