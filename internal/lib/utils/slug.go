// File: internal/lib/utils/slug.go
// Purpose: Generates unique short slugs for URL shortening.

package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateSlug returns a random slug of the given length.
func GenerateSlug(length int) (string, error) {
	slug := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := range slug {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		slug[i] = charset[num.Int64()]
	}

	return string(slug), nil
}
