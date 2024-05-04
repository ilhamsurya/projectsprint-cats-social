package auth

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

func GenerateHash(password, salt []byte) string {
	hashed := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	return base64.StdEncoding.EncodeToString(hashed)
}

func CompareHash(hashedPassword, plainPassword, salt string) error {
	hashSalt := GenerateHash([]byte(plainPassword), []byte(salt))
	if hashedPassword != hashSalt {
		return errors.New("Hash doesn't match")
	}

	return nil
}
