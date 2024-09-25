package validator

import (
	"net/url"
)

func ValidateURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}