![Workflow](https://github.com/eredotpkfr/golidators/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/eredotpkfr/golidators)](https://goreportcard.com/report/github.com/eredotpkfr/golidators)
[![Go Reference](https://pkg.go.dev/badge/github.com/eredotpkfr/golidators.svg)](https://pkg.go.dev/github.com/eredotpkfr/golidators)
[![Go Version](https://img.shields.io/github/go-mod/go-version/eredotpkfr/golidators)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/eredotpkfr/golidators)](https://github.com/eredotpkfr/golidators/releases/latest)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/eredotpkfr/golidators/blob/main/LICENSE)
[![Stars](https://img.shields.io/github/stars/eredotpkfr/golidators?style=social)](https://github.com/eredotpkfr/golidators/stargazers)

# golidators

Golidators is a golang package, it includes basic data validation functions and regexes.

## Install

```bash
~$ go get github.com/eredotpkfr/golidators
```

## Overview

Following validators available on this package:

- Domain
- MD5, SHA1, SHA224, SHA256, SHA512
- IPv4, IPv4CIDR, IPv6, IPv6CIDR
- MAC
- Port
- URL
- UUID
- CreditCard/Luhn

## Usage

Just import and use it. Also see documentation at [pkg.go.dev](https://pkg.go.dev/github.com/eredotpkfr/golidators#section-documentation)

```go
package main

import (
    "github.com/eredotpkfr/golidators"
    "fmt"
)

func main() {
  // Validate domain address
  fmt.Println(golidators.Domain("www.example.com"))
  // Validate IPv4 address
  fmt.Println(golidators.Ipv4("::1"))
  // Validate IPv6 address
  fmt.Println(golidators.Ipv6("::1"))
  // Validate URL
  fmt.Println(golidators.Url("https://www.example.com"))
  // Validate IPv4CIDR
  fmt.Println(golidators.Ipv4Cidr("127.0.0.1/12"))
  // Validate most common hashes
  fmt.Println(golidators.Md5("foo/bar"))
  // Validate with Luhn algorithm
  fmt.Println(golidators.Luhn(5300025108592596))
  // Calculate check digit with Luhn algorithm
  fmt.Println(golidators.LuhnCheckDigit(5146713835433))
  // Validate credit card number with Luhn
  fmt.Println(golidators.CreditCard(5184214431476070))
}
```

## Contact

Blog - [erdoganyoksul.com](https://www.erdoganyoksul.com)<br/>
Mail - erdoganyoksul3@gmail.com
