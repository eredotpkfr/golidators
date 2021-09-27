package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

// Uuid function for validating UUID
func Uuid(uuid string) bool {
	return regexes.UUIDRegex.MatchString(uuid)
}
