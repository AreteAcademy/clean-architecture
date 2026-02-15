package user_test

import (
	"errors"
	"regexp"
	"testing"
	"unicode"
)

// DOMAIN SHARED
var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasLower, hasUpper, hasSpacial bool

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpacial = true
		}
	}

	return hasLower && hasUpper && hasSpacial
}

// INTERFACES

var (
	ErrUserNameIsRequired     = errors.New("name is required")
	ErrUserEmailIsRequired    = errors.New("email is required")
	ErrUserEmailInvalid       = errors.New("email invalid")
	ErrUserPasswordIsRequired = errors.New("password is required")
	ErrUserPasswordInvalid    = errors.New("password invalid")
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

	if user.Email == "" {
		return nil, ErrUserEmailIsRequired
	}

	if !isValidEmail(user.Email) {
		return nil, ErrUserEmailInvalid
	}

	if user.Password == "" {
		return nil, ErrUserPasswordIsRequired
	}

	if !isValidPassword(user.Password) {
		return nil, ErrUserPasswordInvalid
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

func TestCreateUser_ShouldReturnAnErrorIfNameEmpty(t *testing.T) {
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

func TestCreateUser_ShouldReturnAnErrorIfEmailEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	user, err := sut.UseCase.Perform(&User{
		Name:     "Daniel",
		Email:    "",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != ErrUserEmailIsRequired {
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
	user, err := sut.UseCase.Perform(&User{
		Name:     "Daniel",
		Email:    "daniel.com.br",
		Password: "@Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != ErrUserEmailInvalid {
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
	user, err := sut.UseCase.Perform(&User{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != ErrUserPasswordIsRequired {
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
	user, err := sut.UseCase.Perform(&User{
		Name:     "Daniel",
		Email:    "daniel@gmail.com",
		Password: "Daniel123",
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != ErrUserPasswordInvalid {
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
	user, err := sut.UseCase.Perform(&User{
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
}
