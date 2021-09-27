package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

func Uuid(uuid string) bool {
	return regexes.UuidRegex.MatchString(uuid)
}
