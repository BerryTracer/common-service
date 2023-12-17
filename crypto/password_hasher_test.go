package crypto_test

import (
	"testing"

	"github.com/BerryTracer/common-service/crypto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHasher_HashPassword(t *testing.T) {
	hasher := crypto.NewBcryptHasher()
	password := "password"
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Errorf("Hashed password does not match original password: %v", err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestBcryptHasher_HashPassword_InvalidPassword(t *testing.T) {
	hasher := crypto.NewBcryptHasher()
	password := "password"
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("invalid"))
	if err == nil {
		t.Errorf("Hashed password matches invalid password: %v", err)
	}

	assert.Error(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestBcryptHasher_ComparePassword(t *testing.T) {
	hasher := crypto.NewBcryptHasher()
	password := "password"
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = hasher.ComparePassword(password, hashedPassword)
	if err != nil {
		t.Errorf("Hashed password does not match original password: %v", err)
	}

	assert.NoError(t, err)
}

func TestBcryptHasher_ComparePassword_InvalidPassword(t *testing.T) {
	hasher := crypto.NewBcryptHasher()
	password := "password"
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = hasher.ComparePassword("invalid", hashedPassword)
	if err == nil {
		t.Errorf("Hashed password matches invalid password: %v", err)
	}

	assert.Error(t, err)
}
