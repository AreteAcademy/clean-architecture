package user

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
	Repository *GormUserRepository
	DB         *gorm.DB
	User       *domain.User
}

func makeSut(t *testing.T) SUT {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&UserGorm{}))

	repository := NewGoUserRepository(db)

	now := time.Now()

	user := &domain.User{
		ID:        "123456",
		Name:      "Daniel",
		Email:     "daniel@com.br",
		Password:  "@Daniel123",
		CreatedAt: now,
		UpdatedAt: now,
	}

	return SUT{
		Repository: repository,
		DB:         db,
		User:       user,
	}
}

func TestUserRepository_Save_ShouldPersistUser(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	// Act
	require.NoError(t, sut.Repository.Save(sut.User))
	getUser, err := sut.Repository.GetById(sut.User.ID)

	require.NoError(t, err)
	require.NotNil(t, getUser)

	// Assert
	assert.Equal(t, sut.User.ID, getUser.ID)
	assert.Equal(t, sut.User.Name, getUser.Name)
	assert.Equal(t, sut.User.Email, getUser.Email)
	assert.Equal(t, sut.User.Password, getUser.Password)
	assert.Equal(t, sut.User.CreatedAt.Local(), getUser.CreatedAt.Local())
	assert.Equal(t, sut.User.UpdatedAt.Local(), getUser.UpdatedAt.Local())
}

func TestUserRepository_Update_ShouldUpdateUser(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.User))

	expected := struct {
		Name     string
		Email    string
		Password string
	}{
		Name:     "Daniel Editado",
		Email:    "daniel.editado@gmail.com",
		Password: "@@Daniel1234",
	}
	sut.User.Name = expected.Name
	sut.User.Email = expected.Email
	sut.User.Password = expected.Password

	require.NoError(t, sut.Repository.Update(sut.User))
	getUser, err := sut.Repository.GetById(sut.User.ID)

	require.NoError(t, err)
	require.NotNil(t, getUser)

	assert.Equal(t, expected.Name, getUser.Name)
	assert.Equal(t, expected.Email, getUser.Email)
	assert.Equal(t, expected.Password, getUser.Password)
}

func TestUserRepository_Update_ShouldReturnError_WhenUserNotFound(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.User))

	sut.User.ID = "9999"
	err := sut.Repository.Update(sut.User)

	require.Error(t, err)
	assert.ErrorIs(t, err, domain.ErrUserNotFound)
}

func TestUserRepository_Update_ShouldReturnError_WhenUserIsNil(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save(sut.User))
	err := sut.Repository.Update(nil)

	require.Error(t, err)
	assert.ErrorIs(t, err, ErrRepoUserIsNil)
}

func TestUserRepository_GetById_ShouldReturnUser(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save((sut.User)))

	getUser, err := sut.Repository.GetById(sut.User.ID)

	require.NoError(t, err)
	require.NotNil(t, getUser)
	assert.Equal(t, sut.User.ID, getUser.ID)
	assert.Equal(t, sut.User.Name, getUser.Name)
	assert.Equal(t, sut.User.Email, getUser.Email)
	assert.Equal(t, sut.User.Password, getUser.Password)
}

func TestUserRepository_GetById_ShouldReturnNil_WhenUserNotFound(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	require.NoError(t, sut.Repository.Save((sut.User)))
	getUser, err := sut.Repository.GetById("not-found")

	require.Error(t, err)
	require.Nil(t, getUser)
	assert.ErrorIs(t, err, domain.ErrUserNotFound)
}

func TestUserRepository_Count_ShouldReturnCorrectValue(t *testing.T) {
	// Arrange
	sut := makeSut(t)

	user1 := *sut.User
	user1.ID = "User01"
	user1.Name = "User 01"
	user1.Email = "user01@gmail.com"

	user2 := *sut.User
	user2.ID = "User02"
	user2.Name = "User 02"
	user2.Email = "user02@gmail.com"

	user3 := *sut.User
	user3.ID = "User03"
	user3.Name = "User 03"
	user3.Email = "user03@gmail.com"

	require.NoError(t, sut.Repository.Save(&user1))
	require.NoError(t, sut.Repository.Save(&user2))
	require.NoError(t, sut.Repository.Save(&user3))

	count, err := sut.Repository.Count()

	require.NoError(t, err)
	assert.Equal(t, 3, count)
}
