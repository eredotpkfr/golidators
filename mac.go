package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

// Mac function for validating MAC Addresses
func Mac(macAddr string) bool {
	return regexes.MacRegex.MatchString(macAddr)
}
