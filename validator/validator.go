package validator

import (
	"regexp"
)

func ValidateURL(input string) bool {
	var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
    return urlRegex.MatchString(input)
}