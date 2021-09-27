package golidators

import (
	"github.com/eredotpkfr/golidators/internal/regexes"
	"golang.org/x/net/idna"
)

func Domain(domain string) bool {
	var idnap *idna.Profile = idna.New()
	domain, err := idnap.ToASCII(domain)

	if err != nil {
		return false
	}

	return regexes.DomainRegex.MatchString(domain)
}
