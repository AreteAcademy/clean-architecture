package user

import (
	"time"

	"github.com/areteacademy/internal/domain"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type createUserUseCase struct {
	repo   domain.UserRepository
	hasher domain.UserPasswordHasher
}

type CreateUserUseCase interface {
	Perform(input *CreateUserInput) (*CreateUserOutput, error)
}

func NewCreateUserUseCase(repo domain.UserRepository, hasher domain.UserPasswordHasher) CreateUserUseCase {
	return &createUserUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *createUserUseCase) Perform(input *CreateUserInput) (*CreateUserOutput, error) {
	user, err := domain.NewUser(
		input.Name,
		input.Email,
		input.Password,
	)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := uc.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
