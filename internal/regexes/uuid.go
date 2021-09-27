package regexes

import rgx "regexp"

const uuidPattern = "^[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}$"
var UuidRegex = rgx.MustCompile(make_case_insensitive(uuidPattern))
