package auth

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func GenerateHash(password, salt []byte) string {
	hashed := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	return base64.StdEncoding.EncodeToString(hashed)
}
