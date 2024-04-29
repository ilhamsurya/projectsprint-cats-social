package utils

import (
	"crypto/rand"
	"encoding/base64"
	mathrand "math/rand"
	"time"
)

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomAlphaNumeric(n int) string {
	// ascii range (A-Z)
	min := 65
	max := 90

	result := make([]byte, n)
	for i := range result {
		result[i] = byte(mathrand.Intn(max-min+1) + min)
	}

	return string(result)
}
