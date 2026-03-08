package product

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/areteacademy/internal/domain"
	productRepo "github.com/areteacademy/internal/infra/repository/product"
	userRepo "github.com/areteacademy/internal/infra/repository/user"
)

type SUT struct {
	UseCase     ListByUserIdProductUseCase
	ProductRepo *productRepo.InMemoryProductRepository
	UserRepo    *userRepo.InMemoryUserRepository
	Product     *domain.Product
	User        *domain.User
}

func makeSut() SUT {
	productRepo := productRepo.NewInMemoryProductRepository()
	userRepo := userRepo.NewInMemoryUserRepository()
	usecase := NewListByUserIdProductUseCase(productRepo, userRepo)

	now := time.Now()
	user := &domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@gmail.com",
		CreatedAt: now,
		UpdatedAt: now,
	}
	product := &domain.Product{
		ID:          "123456",
		UserId:      "123456",
		CategoryId:  "123456",
		Name:        "Produto1",
		Description: "Meu Produto",
		Status:      "ACTIVE",
		Price:       100,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return SUT{
		UseCase:     usecase,
		ProductRepo: productRepo,
		UserRepo:    userRepo,
		User:        user,
		Product:     product,
	}
}

func TestListByUserIdProduct_ShouldReturnAnError_WhenUserIdEmpty(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)
	sut.ProductRepo.Save(sut.Product)

	// Act
	producties, err := sut.UseCase.Perform("")

	// Assert
	require.Error(t, err)
	require.Nil(t, producties)
	assert.ErrorIs(t, err, domain.ErrProductUserIdIsRequired)
}

func TestListByUserIdProduct_ShouldReturnAnError_WhenUserNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)
	sut.ProductRepo.Save(sut.Product)

	// Act
	producties, err := sut.UseCase.Perform("1234567")

	// Assert
	require.Error(t, err)
	require.Nil(t, producties)
	assert.ErrorIs(t, err, domain.ErrProductUserNotFound)
}

func TestListByUserIdProduct_ShouldReturnAnError_WhenRepoUserFailOnGetById(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)
	sut.UserRepo.FailOnGet = true
	sut.ProductRepo.Save(sut.Product)

	// Act
	producties, err := sut.UseCase.Perform("123456")

	// Assert
	require.Error(t, err)
	require.Nil(t, producties)
	assert.ErrorIs(t, err, userRepo.ErrSimulatedFailureRepoUser)
}

func TestListByUserIdProduct_ShouldReturnAnError_WhenProductNotFound(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)
	sut.Product.UserId = "1234567"
	sut.ProductRepo.Save(sut.Product)

	// Act
	producties, err := sut.UseCase.Perform("123456")

	// Assert
	require.Error(t, err)
	require.Nil(t, producties)
	assert.ErrorIs(t, err, domain.ErrProductNotFound)
}

func TestListByUserIdProduct_ShouldReturnAnError_WhenRepoProductFailOnListByUserId(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)
	sut.ProductRepo.Save(sut.Product)
	sut.ProductRepo.FailOnList = true

	// Act
	producties, err := sut.UseCase.Perform("123456")

	// Assert
	require.Error(t, err)
	require.Nil(t, producties)
	assert.ErrorIs(t, err, productRepo.ErrSimulatedFailureRepoProduct)
}

func TestListByUserIdProduct_ShouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()
	sut.UserRepo.Save(sut.User)

	p01 := *sut.Product // Copy
	p01.ID = "p01"
	sut.ProductRepo.Save(&p01)

	p02 := *sut.Product // Copy
	p02.ID = "p02"
	sut.ProductRepo.Save(&p02)

	p03 := *sut.Product // Copy
	p03.ID = "p03"
	sut.ProductRepo.Save(&p03)

	expected := ListByUserIdProductOutput{
		{
			ID:          p01.ID,
			UserId:      p01.UserId,
			CategoryId:  p01.CategoryId,
			Name:        p01.Name,
			Description: p01.Description,
			Status:      p01.Status,
			Price:       p01.Price,
			CreatedAt:   p01.CreatedAt,
			UpdatedAt:   p01.UpdatedAt,
		},
		{
			ID:          p02.ID,
			UserId:      p02.UserId,
			CategoryId:  p02.CategoryId,
			Name:        p02.Name,
			Description: p02.Description,
			Status:      p02.Status,
			Price:       p02.Price,
			CreatedAt:   p02.CreatedAt,
			UpdatedAt:   p02.UpdatedAt,
		},
		{
			ID:          p03.ID,
			UserId:      p03.UserId,
			CategoryId:  p03.CategoryId,
			Name:        p03.Name,
			Description: p03.Description,
			Status:      p03.Status,
			Price:       p03.Price,
			CreatedAt:   p03.CreatedAt,
			UpdatedAt:   p03.UpdatedAt,
		},
	}

	// Act
	producties, err := sut.UseCase.Perform("123456")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, producties)

	assert.ElementsMatch(t, expected, producties)
}
