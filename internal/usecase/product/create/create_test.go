package product

import (
	"testing"

	"github.com/areteacademy/internal/domain"
	categoryRepo "github.com/areteacademy/internal/infra/repository/category"
	productRepo "github.com/areteacademy/internal/infra/repository/product"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase      CreateProductUseCase
	ProductRepo  *productRepo.InMemoryProductRepository
	CategoryRepo *categoryRepo.InMemoryCategoryRepository
	UserRepo     *userRepo.InMemoryUserRepository
}

func makeSut() SUT {
	productRepo := productRepo.NewInMemoryProductRepository()
	categoryRepo := categoryRepo.NewInMemoryCategoryRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewCreateProductUseCase(productRepo, categoryRepo, userRepo)

	return SUT{
		UseCase:      usecase,
		ProductRepo:  productRepo,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

func TestCreateProduct_ShouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	product, err := sut.UseCase.Perform(CreateProductInput{
		UserId:      "",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu produto",
		Status:      "ACTIVE",
		Price:       100,
	})

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err != domain.ErrProductUserIdIsRequired {
		t.Errorf("expected ErrProductUserIdIsRequired, got %v", err)
	}

	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
}
