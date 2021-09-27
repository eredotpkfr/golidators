package regexes

import rgx "regexp"

const domainPattern string = "^(?:[a-zA-Z0-9]" +
  "(?:[a-zA-Z0-9-_]{0,61}[A-Za-z0-9])?\\.)" +
  "+[A-Za-z0-9][A-Za-z0-9-_]{0,61}" +
  "[A-Za-z]$"

var DomainRegex = rgx.MustCompile(make_case_insensitive(domainPattern))
