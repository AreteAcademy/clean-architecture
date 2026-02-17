package user

import "github.com/areteacademy/internal/domain"

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
	if id == "" {
		return nil, domain.ErrUserIdIsRequired
	}

	user, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
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
