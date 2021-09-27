package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

func Url(url string) bool {
	return regexes.UrlRegex.MatchString(url)
}
