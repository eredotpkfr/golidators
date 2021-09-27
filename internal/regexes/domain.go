package regexes

import rgx "regexp"

const domainPattern string = "^(?:[a-zA-Z0-9]" +
	"(?:[a-zA-Z0-9-_]{0,61}[A-Za-z0-9])?\\.)" +
	"+[A-Za-z0-9][A-Za-z0-9-_]{0,61}" +
	"[A-Za-z]$"

// DomainRegex for validating domains
var DomainRegex = rgx.MustCompile(makeCaseInsensitive(domainPattern))
