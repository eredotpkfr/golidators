package golidators

import "github.com/eredotpkfr/golidators/internal/regexes"

func Md5(hash string) bool {
	return regexes.Md5Regex.MatchString(hash)
}

func Sha1(hash string) bool {
	return regexes.Sha1Regex.MatchString(hash)
}

func Sha224(hash string) bool {
	return regexes.Sha224Regex.MatchString(hash)
}

func Sha256(hash string) bool {
	return regexes.Sha256Regex.MatchString(hash)
}

func Sha512(hash string) bool {
	return regexes.Sha512Regex.MatchString(hash)
}
