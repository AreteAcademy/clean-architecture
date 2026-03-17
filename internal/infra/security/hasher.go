package user

import (
	"github.com/areteacademy/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordhasher struct{}

func NewBcryptPasswordHasher() *BcryptPasswordhasher {
	return &BcryptPasswordhasher{}
}

func (h *BcryptPasswordhasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

var _ domain.UserPasswordHasher = (*BcryptPasswordhasher)(nil)
