package user

import (
	"testing"
	"time"

	"github.com/areteacademy/internal/domain"
	repo "github.com/areteacademy/internal/infra/repository/user"
	security "github.com/areteacademy/internal/infra/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SUT struct {
	UseCase CreateUserUseCase
	Repo    *repo.InMemoryUserRepository
	User    *domain.User
}

func makeSut() SUT {
	repo := repo.NewInMemoryUserRepository()
	hash := security.NewBcryptPasswordHasher()
	usecase := NewCreateUserUseCase(repo, hash)

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
		UseCase: usecase,
		Repo:    repo,
		User:    user,
	}
}

func validInput(sut SUT) CreateUserInput {
	return CreateUserInput{
		Name:     sut.User.Name,
		Email:    sut.User.Email,
		Password: sut.User.Password,
	}
}

func TestCreateUser_GivenInvalidInput_ShouldReturnError(t *testing.T) {
	// Act
	testCases := []struct {
		name        string
		setup       func(sut SUT)
		input       func(sut SUT) CreateUserInput
		expectedErr error
	}{
		{
			name:  "Empty Name",
			setup: func(sut SUT) {},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				in.Name = ""
				return in
			},
			expectedErr: domain.ErrUserNameIsRequired,
		},
		{
			name:  "Empty Email",
			setup: func(sut SUT) {},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				in.Email = ""
				return in
			},
			expectedErr: domain.ErrUserEmailIsRequired,
		},
		{
			name:  "Invalid Email",
			setup: func(sut SUT) {},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				in.Email = "daniel.com.br"
				return in
			},
			expectedErr: domain.ErrUserEmailInvalid,
		},
		{
			name:  "Empty Password",
			setup: func(sut SUT) {},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				in.Password = ""
				return in
			},
			expectedErr: domain.ErrUserPasswordIsRequired,
		},
		{
			name:  "Invalid Password",
			setup: func(sut SUT) {},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				in.Password = "Daniel123"
				return in
			},
			expectedErr: domain.ErrUserPasswordInvalid,
		},
		{
			name: "Repo User Fail On Save",
			setup: func(sut SUT) {
				sut.Repo.FailOnSave = true
			},
			input: func(sut SUT) CreateUserInput {
				in := validInput(sut)
				return in
			},
			expectedErr: repo.ErrSimulatedFailureRepoUser,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			sut := makeSut()

			if tc.setup != nil {
				tc.setup(sut)
			}

			input := tc.input(sut)

			user, err := sut.UseCase.Perform(&input)

			// Assert
			require.Error(t, err)
			require.Nil(t, user)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestCreateUser_shouldReturnAnError_WhenValidationFails(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	_, _ = sut.UseCase.Perform(&CreateUserInput{
		Name:     "",
		Email:    "daniel@gmail.com",
		Password: "@Danel123",
	})

	// Assert
	count, err := sut.Repo.Count()
	require.Nil(t, err)
	require.Equal(t, count, 0)
}

func TestCreateUser_shouldReturnSuccess(t *testing.T) {
	// Arrange
	sut := makeSut()

	// Act
	expected := &CreateUserInput{
		Name:     sut.User.Name,
		Email:    sut.User.Email,
		Password: sut.User.Password,
	}
	user, err := sut.UseCase.Perform(expected)

	// Assert
	require.Nil(t, err)
	require.NotNil(t, user)

	assert.NotNil(t, user.ID)

	assert.Equal(t, expected.Name, user.Name)
	assert.Equal(t, expected.Email, user.Email)

	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())

	count, err := sut.Repo.Count()
	assert.Equal(t, count, 1)
}
