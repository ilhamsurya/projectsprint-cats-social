package utils

import (
	"regexp"
)

func IsValidFullName(fullname string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]{3,40}$`)
	return re.MatchString(fullname)
}
