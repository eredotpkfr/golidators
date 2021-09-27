package regexes

import rgx "regexp"

const (
  md5Pattern string = "^[0-9a-f]{32}$"
  sha1Pattern string = "^[0-9a-f]{40}$"
  sha224Pattern string = "^[0-9a-f]{56}$"
  sha256Pattern string = "^[0-9a-f]{64}$"
  sha512Pattern string = "^[0-9a-f]{128}$"
)

var (
  Md5Regex = rgx.MustCompile(make_case_insensitive(md5Pattern))
  Sha1Regex = rgx.MustCompile(make_case_insensitive(sha1Pattern))
  Sha224Regex = rgx.MustCompile(make_case_insensitive(sha224Pattern))
  Sha256Regex = rgx.MustCompile(make_case_insensitive(sha256Pattern))
  Sha512Regex = rgx.MustCompile(make_case_insensitive(sha512Pattern))
)
