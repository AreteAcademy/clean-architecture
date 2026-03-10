package category

import (
	"errors"

	"github.com/areteacademy/internal/domain"
)

var ErrSimulatedFailureRepoCategory = errors.New("database error")

type InMemoryCategoryRepository struct {
	FailOnSave             bool
	FailOnUpdate           bool
	FailOnGetById          bool
	FailOnGetByIdAndUserId bool
	FailOnList             bool
	FailOnCount            bool
	categories             map[string]*domain.Category
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		categories: make(map[string]*domain.Category),
	}
}

func (r *InMemoryCategoryRepository) Save(category *domain.Category) error {
	if r.FailOnSave {
		return ErrSimulatedFailureRepoCategory
	}
	r.categories[category.ID] = category
	return nil
}

func (r *InMemoryCategoryRepository) Update(category *domain.Category) error {
	if r.FailOnUpdate {
		return ErrSimulatedFailureRepoCategory
	}
	r.categories[category.ID] = category
	return nil
}

func (r *InMemoryCategoryRepository) GetById(id string) (*domain.Category, error) {
	if r.FailOnGetById {
		return nil, ErrSimulatedFailureRepoCategory
	}
	category, exists := r.categories[id]
	if !exists {
		return nil, nil
	}
	return category, nil
}

func (r *InMemoryCategoryRepository) GetByIdAndUserId(id, userId string) (*domain.Category, error) {
	if r.FailOnGetByIdAndUserId {
		return nil, ErrSimulatedFailureRepoCategory
	}

	for _, c := range r.categories {
		if c.UserId == userId && c.ID == id {
			return c, nil
		}
	}

	return nil, nil
}

func (r *InMemoryCategoryRepository) ListByUserId(userId string) ([]*domain.Category, error) {
	if r.FailOnList {
		return nil, ErrSimulatedFailureRepoCategory
	}

	var categories []*domain.Category
	for _, c := range r.categories {
		if c.UserId == userId {
			categories = append(categories, c)
		}
	}

	return categories, nil
}

func (r *InMemoryCategoryRepository) Count() (int, error) {
	if r.FailOnCount {
		return 0, ErrSimulatedFailureRepoCategory
	}
	return len(r.categories), nil
}

var _ domain.CategoryRepository = (*InMemoryCategoryRepository)(nil)
