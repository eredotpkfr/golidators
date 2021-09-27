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
  // Md5Regex for validating MD5
  Md5Regex = rgx.MustCompile(makeCaseInsensitive(md5Pattern))
  // Sha1Regex for validating SHA1
  Sha1Regex = rgx.MustCompile(makeCaseInsensitive(sha1Pattern))
  // Sha224Regex for validating SHA224
  Sha224Regex = rgx.MustCompile(makeCaseInsensitive(sha224Pattern))
  // Sha256Regex for validating SHA256
  Sha256Regex = rgx.MustCompile(makeCaseInsensitive(sha256Pattern))
  // Sha512Regex for validating SHA512
  Sha512Regex = rgx.MustCompile(makeCaseInsensitive(sha512Pattern))
)
