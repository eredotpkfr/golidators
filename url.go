package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

// Url function for validating URL's
func Url(url string) bool {
	return regexes.URLRegex.MatchString(url)
}
