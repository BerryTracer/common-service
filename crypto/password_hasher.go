package crypto

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password string, hashedPassword string) error
}

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func (h *BcryptHasher) ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
