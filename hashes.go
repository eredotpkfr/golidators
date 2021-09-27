package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

// Md5 function for validating MD5
func Md5(hash string) bool {
	return regexes.Md5Regex.MatchString(hash)
}

// Sha1 function for validating SHA1
func Sha1(hash string) bool {
	return regexes.Sha1Regex.MatchString(hash)
}

// Sha224 function for validating SHA224
func Sha224(hash string) bool {
	return regexes.Sha224Regex.MatchString(hash)
}

// Sha256 function for validating SHA256
func Sha256(hash string) bool {
	return regexes.Sha256Regex.MatchString(hash)
}

// Sha512 function for validating SHA512
func Sha512(hash string) bool {
	return regexes.Sha512Regex.MatchString(hash)
}
