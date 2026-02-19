package user

import "github.com/areteacademy/internal/domain"

type updateUserUseCase struct {
	repo domain.UserRepository
}

type UpdateUserUseCase interface {
	Perform(input UpdateUserInput) (*UpdateUserOutput, error)
}

func NewUpdateUserUseCase(repo domain.UserRepository) UpdateUserUseCase {
	return &updateUserUseCase{
		repo: repo,
	}
}

func (uc *updateUserUseCase) Perform(input UpdateUserInput) (*UpdateUserOutput, error) {
	user, err := domain.Update(
		input.ID,
		input.Name,
		input.Email,
	)
	if err != nil {
		return nil, err
	}

	exists, err := uc.repo.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, domain.ErrUserNotFound
	}

	if err := uc.repo.Update(user); err != nil {
		return nil, err
	}

	return &UpdateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: exists.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
