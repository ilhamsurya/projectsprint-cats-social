package utils

import (
	"unicode"
)

const minPasswordLen = 5
const maxPasswordLen = 15

func IsSolidPassword(s string) bool {
	var (
		hasMinMaxLen = false
		hasNumber    = false
		hasLetter    = false
	)

	if len(s) >= minPasswordLen && len(s) <= maxPasswordLen {
		hasMinMaxLen = true
	}

	for _, char := range s {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	return hasMinMaxLen && (hasLetter || hasNumber)
}
