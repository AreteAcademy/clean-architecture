package category

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
	Repository *GormCategoryRepository
	DB         *gorm.DB
	Category   *domain.Category
}

func makeSut(t *testing.T) SUT {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&CategoryGorm{}))

	repository := NewGormCategoryRepository(db)

	now := time.Now()

	category := &domain.Category{
		ID:        "cat-01",
		UserId:    "user-01",
		Name:      "Categoria",
		Status:    string(domain.CategoryStatusActive),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return SUT{
		Repository: repository,
		DB:         db,
		Category:   category,
	}
}

func TestCategoryRepository_Save_ShouldPersistCategory(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))
	getCategory, err := sut.Repository.GetById(sut.Category.ID)

	require.NoError(t, err)
	require.NotNil(t, getCategory)
	assert.Equal(t, sut.Category.ID, getCategory.ID)
	assert.Equal(t, sut.Category.UserId, getCategory.UserId)
	assert.Equal(t, sut.Category.Name, getCategory.Name)
	assert.Equal(t, sut.Category.Status, getCategory.Status)
	assert.Equal(t, sut.Category.CreatedAt.Local(), getCategory.CreatedAt.Local())
	assert.Equal(t, sut.Category.UpdatedAt.Local(), getCategory.UpdatedAt.Local())
}

func TestCategoryRepository_Save_ShouldReturnError_WhenCategoryIsNil(t *testing.T) {
	sut := makeSut(t)

	err := sut.Repository.Save(nil)

	require.Error(t, err)
	assert.ErrorIs(t, err, ErrRepositoryCategoryNil)
}

func TestCategoryRepository_Update_ShouldUpdateCategory(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))

	expected := struct {
		Name   string
		Status string
	}{
		Name:   "Categoria editada",
		Status: string(domain.CategoryStatusInactive),
	}

	sut.Category.Name = expected.Name
	sut.Category.Status = expected.Status

	require.NoError(t, sut.Repository.Update(sut.Category))
	getCategory, err := sut.Repository.GetById(sut.Category.ID)

	require.NoError(t, err)
	require.NotNil(t, getCategory)
	assert.Equal(t, expected.Name, getCategory.Name)
	assert.Equal(t, expected.Status, getCategory.Status)
}

func TestCategoryRepository_Update_ShoukdReturnError_WhenCategoryNotFond(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))

	sut.Category.ID = "9999"
	err := sut.Repository.Update(sut.Category)

	require.Error(t, err)
	assert.ErrorIs(t, err, domain.ErrCategoryNotFound)
}

func TestCategoryRepository_Update_ShouldReturnError_WhenCategoryIsNil(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))
	err := sut.Repository.Update(nil)

	require.Error(t, err)
	assert.ErrorIs(t, err, ErrRepositoryCategoryNil)
}

func TestCategoryRepository_GetById_ShouldReturnCategory(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))

	getCategory, err := sut.Repository.GetById(sut.Category.ID)

	require.NoError(t, err)
	require.NotNil(t, getCategory)
	assert.Equal(t, sut.Category.ID, getCategory.ID)
	assert.Equal(t, sut.Category.Name, getCategory.Name)
	assert.Equal(t, sut.Category.Status, getCategory.Status)
}

func TestCategoryRepository_GetById_ShouldReturnNil_WhenCategoryNotFound(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))
	getCategory, err := sut.Repository.GetById("not-found")

	require.Error(t, err)
	require.Nil(t, getCategory)
	assert.ErrorIs(t, err, domain.ErrCategoryNotFound)
}

func TestCategoryRepository_GetByIdAndUserId_ShouldReturnCategory(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))
	getCategory, err := sut.Repository.GetByIdAndUserId(sut.Category.ID, sut.Category.UserId)

	require.NoError(t, err)
	require.NotNil(t, getCategory)
	assert.Equal(t, sut.Category.ID, getCategory.ID)
	assert.Equal(t, sut.Category.UserId, getCategory.UserId)
	assert.Equal(t, sut.Category.Name, getCategory.Name)
	assert.Equal(t, sut.Category.Status, getCategory.Status)
}

func TestCategoryRepository_GetByIdAndUserId_ShouldReturnNil_WhenCategoryNotFound(t *testing.T) {
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.Category))
	getCategory, err := sut.Repository.GetByIdAndUserId("id-not-found", "user-id-not-found")

	require.Error(t, err)
	require.Nil(t, getCategory)
	assert.ErrorIs(t, err, domain.ErrCategoryNotFound)
}

func TestCategoryRepository_ListByUserId_ShouldReturnCategories(t *testing.T) {
	sut := makeSut(t)

	category1 := *sut.Category
	category1.ID = "cat01"
	category1.Name = "cat 01"

	category2 := *sut.Category
	category2.ID = "cat02"
	category2.Name = "cat 02"

	category3 := *sut.Category
	category3.ID = "cat03"
	category3.Name = "cat 03"

	category4 := *sut.Category
	category4.ID = "cat04"
	category4.UserId = "user-9999"
	category4.Name = "cat 04"

	require.NoError(t, sut.Repository.Save(&category1))
	require.NoError(t, sut.Repository.Save(&category2))
	require.NoError(t, sut.Repository.Save(&category3))
	require.NoError(t, sut.Repository.Save(&category4))

	categories, err := sut.Repository.ListByUserId(sut.Category.UserId)

	require.NoError(t, err)
	require.NotNil(t, categories)
	assert.Len(t, categories, 3)
	assert.ElementsMatch(
		t,
		[]string{"cat01", "cat02", "cat03"},
		[]string{categories[0].ID, categories[1].ID, categories[2].ID},
	)
}

func TestCategoryRepository_Count_ShouldReturnCorrectValue(t *testing.T) {
	sut := makeSut(t)

	category1 := *sut.Category
	category1.ID = "cat01"
	category1.Name = "cat 01"

	category2 := *sut.Category
	category2.ID = "cat02"
	category2.Name = "cat 02"

	category3 := *sut.Category
	category3.ID = "cat03"
	category3.Name = "cat 03"

	require.NoError(t, sut.Repository.Save(&category1))
	require.NoError(t, sut.Repository.Save(&category2))
	require.NoError(t, sut.Repository.Save(&category3))

	count, err := sut.Repository.Count()

	require.NoError(t, err)
	assert.Equal(t, 3, count)
}
