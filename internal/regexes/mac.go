package regexes

import rgx "regexp"

const macPattern string = "^(?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2}$"
var MacRegex = rgx.MustCompile(make_case_insensitive(macPattern))
