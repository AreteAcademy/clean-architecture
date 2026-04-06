package product

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SUT struct {
	Repository *GormProductRepository
	DB         *gorm.DB
	Product    *domain.Product
}

func makeSut(t *testing.T) SUT {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&ProductGorm{}))

	repository := NewGormProductRepository(db)

	now := time.Now()

	product := &domain.Product{
		ID:          "product-123",
		UserId:      "user-123",
		CategoryId:  "category-123",
		Name:        "Notebook",
		Description: "Notebook para dev",
		Status:      string(domain.ProductStatusActive),
		Price:       5000,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return SUT{
		Repository: repository,
		DB:         db,
		Product:    product,
	}
}

func TestProductRepository_Save_ShouldPersistProduct(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Product))

	getProduct, err := sut.Repository.GetById(sut.Product.ID)

	require.NoError(t, err)
	require.NotNil(t, getProduct)

	assert.Equal(t, sut.Product.ID, getProduct.ID)
	assert.Equal(t, sut.Product.UserId, getProduct.UserId)
	assert.Equal(t, sut.Product.CategoryId, getProduct.CategoryId)
	assert.Equal(t, sut.Product.Name, getProduct.Name)
	assert.Equal(t, sut.Product.Description, getProduct.Description)
	assert.Equal(t, sut.Product.Status, getProduct.Status)
	assert.Equal(t, sut.Product.Price, getProduct.Price)
	assert.Equal(t, sut.Product.CreatedAt.Local(), getProduct.CreatedAt.Local())
	assert.Equal(t, sut.Product.UpdatedAt.Local(), getProduct.UpdatedAt.Local())
}

func TestProductRepository_Update_ShouldReturnError_WhenProductNotFound(t *testing.T) {
	sut := makeSut(t)

	sut.Product.ID = "product-123456"

	err := sut.Repository.Update(sut.Product)

	require.Error(t, err)
	assert.ErrorIs(t, err, domain.ErrProductNotFound)
}

func TestProductRepository_Update_ShouldUpdateProduct(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Product))

	sut.Product.CategoryId = "category-updated"
	sut.Product.Name = "Notebook updated"
	sut.Product.Price = 1200

	require.NoError(t, sut.Repository.Update(sut.Product))

	getProduct, err := sut.Repository.GetById(sut.Product.ID)

	require.NoError(t, err)
	require.NotNil(t, getProduct)

	assert.Equal(t, sut.Product.CategoryId, getProduct.CategoryId)
	assert.Equal(t, sut.Product.Name, getProduct.Name)
	assert.Equal(t, sut.Product.Price, getProduct.Price)
}

func TestProductRepository_GetById_ShouldReturnProduct(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Product))

	getProduct, err := sut.Repository.GetById(sut.Product.ID)

	require.NoError(t, err)
	require.NotNil(t, getProduct)

	assert.Equal(t, sut.Product.ID, getProduct.ID)
	assert.Equal(t, sut.Product.UserId, getProduct.UserId)
	assert.Equal(t, sut.Product.CategoryId, getProduct.CategoryId)
}

func TestProductRepository_GetByIdAndUserId_ShouldReturnProduct(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Product))

	getProduct, err := sut.Repository.GetByIdAndUserId(sut.Product.ID, sut.Product.UserId)

	require.NoError(t, err)
	require.NotNil(t, getProduct)

	assert.Equal(t, sut.Product.ID, getProduct.ID)
	assert.Equal(t, sut.Product.UserId, getProduct.UserId)
	assert.Equal(t, sut.Product.CategoryId, getProduct.CategoryId)
}

func TestProductRepository_ListByIdAndUserId_ShouldReturnProduct(t *testing.T) {
	sut := makeSut(t)

	product1 := *sut.Product
	product1.ID = "product-01"
	product1.Name = "product-name-01"

	product2 := *sut.Product
	product2.ID = "product-02"
	product2.Name = "product-name-02"

	product3 := *sut.Product
	product3.ID = "product-03"
	product3.Name = "product-name-03"

	require.NoError(t, sut.Repository.Save(&product1))
	require.NoError(t, sut.Repository.Save(&product2))
	require.NoError(t, sut.Repository.Save(&product3))

	products, err := sut.Repository.ListByUserId(sut.Product.UserId)

	require.NoError(t, err)
	require.NotNil(t, products)

	ids := []string{products[0].ID, products[1].ID, products[2].ID}
	assert.ElementsMatch(t, []string{product1.ID, product2.ID, product3.ID}, ids)
}

func TestProductRepository_Count_ShouldReturnCorrectValue(t *testing.T) {
	sut := makeSut(t)

	product1 := *sut.Product
	product1.ID = "product-01"
	product1.Name = "product-name-01"

	product2 := *sut.Product
	product2.ID = "product-02"
	product2.Name = "product-name-02"

	product3 := *sut.Product
	product3.ID = "product-03"
	product3.Name = "product-name-03"

	require.NoError(t, sut.Repository.Save(&product1))
	require.NoError(t, sut.Repository.Save(&product2))
	require.NoError(t, sut.Repository.Save(&product3))

	count, err := sut.Repository.Count()

	require.NoError(t, err)
	assert.Equal(t, 3, count)
}
