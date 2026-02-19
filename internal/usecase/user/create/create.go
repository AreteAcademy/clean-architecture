package user

import "github.com/areteacademy/internal/domain"

type createUserUseCase struct {
	repo domain.UserRepository
}

type CreateUserUseCase interface {
	Perform(input *CreateUserInput) (*CreateUserOutput, error)
}

func NewCreateUserUseCase(repo domain.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		repo: repo,
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
