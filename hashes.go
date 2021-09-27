package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

// Md5 function for validating MD5
func Md5(md5 string) bool {
	return regexes.Md5Regex.MatchString(md5)
}

// Sha1 function for validating SHA1
func Sha1(sha1 string) bool {
	return regexes.Sha1Regex.MatchString(sha1)
}

// Sha224 function for validating SHA224
func Sha224(sha224 string) bool {
	return regexes.Sha224Regex.MatchString(sha224)
}

// Sha256 function for validating SHA256
func Sha256(sha256 string) bool {
	return regexes.Sha256Regex.MatchString(sha256)
}

// Sha512 function for validating SHA512
func Sha512(sha512 string) bool {
	return regexes.Sha512Regex.MatchString(sha512)
}
