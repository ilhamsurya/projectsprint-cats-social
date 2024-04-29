package utils

import "regexp"

func IsValidPhoneNumber(phoneNum string) bool {
	isValidPhoneNumber := regexp.MustCompile(`^[0-9]{10,13}$`).MatchString
	return isValidPhoneNumber(phoneNum)
}
