package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

func Mac(mac_addr string) bool {
	return regexes.MacRegex.MatchString(mac_addr)
}
