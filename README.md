# golidators
Golidators is a golang package, it includes basic data validation functions and regexes.

## Install
```bash
go get github.com/eredotpkfr/golidators
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

## Usage
Just import and use it
```go
package main

import (
    "github.com/eredotpkfr/golidators"
    "fmt"
)

func main() {
  fmt.Println(goliators.Domain("www.example.com"))
  // true
  fmt.Println(golidators.Ipv4("::1"))
  // false
  fmt.Println(golidators.Ipv6("::1"))
  // true
  fmt.Println(golidators.URL("https://www.example.com"))
  // true
  fmt.Println(golidators.Ipv4Cidr("127.0.0.1/12"))
  // true
  fmt.Println(golidators.Md5("foo/bar"))
  // false
}
```
## Contact
Blog - [erdoganyoksul.com](https://www.erdoganyoksul.com)<br/>
Mail - erdoganyoksul3@gmail.com
