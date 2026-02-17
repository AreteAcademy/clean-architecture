package user

import (
	"errors"

	"github.com/areteacademy/internal/domain"
)

var ErrSimulatedFailureRepoUser = errors.New("database error")

type InMemoryUserRepository struct {
	FailOnSave  bool
	FailOnGet   bool
	FailOnCount bool
	users       map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	if r.FailOnSave {
		return ErrSimulatedFailureRepoUser
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetById(id string) (*domain.User, error) {
	if r.FailOnGet {
		return nil, ErrSimulatedFailureRepoUser
	}
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) Count() (int, error) {
	if r.FailOnCount {
		return 0, ErrSimulatedFailureRepoUser
	}
	return len(r.users), nil
}

var _ domain.UserRepository = (*InMemoryUserRepository)(nil)
