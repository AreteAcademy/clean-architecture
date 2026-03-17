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
	assert.Equal(t, sut.User.CreatedAt.Local(), getUser.CreatedAt.Local())
	assert.Equal(t, sut.User.UpdatedAt.Local(), getUser.UpdatedAt.Local())
}
