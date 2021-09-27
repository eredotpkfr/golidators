package regexes

import rgx "regexp"

const uuidPattern = "^[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}$"

// UUIDRegex for validating UUID data
var UUIDRegex = rgx.MustCompile(makeCaseInsensitive(uuidPattern))
