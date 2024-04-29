package utils

import (
	"projectsphere/cats-social/pkg/protocol/msg"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

func IsPasswordContainUsername(username string, password string) error {
	if password == "" {
		return errors.New(msg.ErrFieldPasswordEmpty)
	}

	if strings.Contains(password, username) {
		return errors.New(msg.ErrPasswordContainUsername)
	}

	return nil
}

func IsSolidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(s) >= 6 {
		hasMinLen = true
	}

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
