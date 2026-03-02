package product

import (
	"errors"

	"github.com/areteacademy/internal/domain"
)

var ErrSimulatedFailureRepoProduct = errors.New("database error")

type InMemoryProductRepository struct {
	FailOnSave   bool
	FailOnUpdate bool
	FailOnGet    bool
	FailOnList   bool
	FailOnCount  bool
	producties   map[string]*domain.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		producties: make(map[string]*domain.Product),
	}
}

func (r *InMemoryProductRepository) Save(product *domain.Product) error {
	if r.FailOnSave {
		return ErrSimulatedFailureRepoProduct
	}
	r.producties[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) Update(product *domain.Product) error {
	if r.FailOnUpdate {
		return ErrSimulatedFailureRepoProduct
	}
	r.producties[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) GetById(id string) (*domain.Product, error) {
	if r.FailOnGet {
		return nil, ErrSimulatedFailureRepoProduct
	}
	product, exists := r.producties[id]
	if !exists {
		return nil, nil
	}
	return product, nil
}

func (r *InMemoryProductRepository) ListByUserId(userId string) ([]*domain.Product, error) {
	if r.FailOnList {
		return nil, ErrSimulatedFailureRepoProduct
	}

	var producties []*domain.Product
	for _, c := range r.producties {
		if c.UserId == userId {
			producties = append(producties, c)
		}
	}

	return producties, nil
}

func (r *InMemoryProductRepository) Count() (int, error) {
	if r.FailOnCount {
		return 0, ErrSimulatedFailureRepoProduct
	}
	return len(r.producties), nil
}

var _ domain.ProductRepository = (*InMemoryProductRepository)(nil)
