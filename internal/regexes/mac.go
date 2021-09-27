package regexes

import rgx "regexp"

const macPattern string = "^(?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2}$"
// MacRegex for validating MAC Addresses
var MacRegex = rgx.MustCompile(makeCaseInsensitive(macPattern))
