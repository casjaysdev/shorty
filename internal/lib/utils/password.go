// File: internal/lib/utils/password.go
// Purpose: Provides secure password hashing and verification utilities.

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a plain password with its hashed version.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
