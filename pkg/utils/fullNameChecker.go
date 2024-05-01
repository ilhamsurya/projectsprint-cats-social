package utils

import (
	"regexp"
)

func IsValidFullName(fullname string) bool {
	// contains only letters and spaces, and must be between 3 and 40 characters long
	re := regexp.MustCompile(`^[a-zA-Z\s]{5,15}$`)
	return re.MatchString(fullname)
}
