package golidators

import (
	"strings"

	utils "github.com/eredotpkfr/golidators/internal/utilities"
)

const (
	// IP CIDR separator
	cidrSep = "/"

	// 0xFF -> 255
	ipv4MaxPart = 0xFF
	ipv4Length  = 4

	// 0xFFFF -> 65535
	ipv6MaxPart = 0xFFFF
	ipv6Length  = 16
)

// Ipv4 function for validating IPv4
func Ipv4(ipv4Addr string) bool {
	for index := 0; index < ipv4Length; index++ {

		if len(ipv4Addr) == 0 {
			return false
		}

		if index > 0 {
			if ipv4Addr[0] != '.' {
				return false
			}

			ipv4Addr = ipv4Addr[1:]
		}

		number, consumed, ok := utils.Dtoi(ipv4Addr)

		if !ok || number > ipv4MaxPart {
			return false
		}

		if consumed > 1 && ipv4Addr[0] == '0' {
			return false
		}

		ipv4Addr = ipv4Addr[consumed:]

	}

	return len(ipv4Addr) == 0
}

// Ipv4Cidr function for validating IPv4CIDR
func Ipv4Cidr(ipv4AddrCidr string) bool {
	splitted := strings.Split(ipv4AddrCidr, cidrSep)

	if len(splitted) != 2 {
		return false
	}

	prefix, suffix := splitted[0], splitted[1]
	number, consumed, ok := utils.Dtoi(suffix)

	if !ok || !Ipv4(prefix) || consumed > 2 {
		return false
	}

	return 0 <= number && number <= 32
}

// Ipv6 function for validating IPv6
func Ipv6(ipv6Addr string) bool {
	ellipsis := -1

	if len(ipv6Addr) >= 2 && ipv6Addr[0] == ':' && ipv6Addr[1] == ':' {
		ellipsis = 0
		ipv6Addr = ipv6Addr[2:]

		if len(ipv6Addr) == 0 {
			return true
		}
	}

	index := 0
	for index < ipv6Length {
		number, consumed, ok := utils.Xtoi(ipv6Addr)

		if !ok || number > ipv6MaxPart {
			return false
		}

		if consumed < len(ipv6Addr) && ipv6Addr[consumed] == '.' {
			if ellipsis < 0 && index != ipv6Length-ipv4Length {
				return false
			}
			if index+ipv4Length > ipv6Length {
				return false
			}
			if !Ipv4(ipv6Addr) {
				return false
			}

			ipv6Addr = ""
			index += ipv4Length
			break
		}

		index += 2

		ipv6Addr = ipv6Addr[consumed:]
		if len(ipv6Addr) == 0 {
			break
		}

		if ipv6Addr[0] != ':' || len(ipv6Addr) == 1 {
			return false
		}

		ipv6Addr = ipv6Addr[1:]

		if ipv6Addr[0] == ':' {
			if ellipsis >= 0 {
				return false
			}

			ellipsis = index
			ipv6Addr = ipv6Addr[1:]

			if len(ipv6Addr) == 0 {
				break
			}

		}
	}

	if len(ipv6Addr) != 0 {
		return false
	}

	if index < ipv6Length {
		if ellipsis < 0 {
			return false
		}
	} else if ellipsis >= 0 {
		return false
	}

	return true
}

// Ipv6Cidr function for validating IPv6CIDR
func Ipv6Cidr(ipv6AddrCidr string) bool {
	splitted := strings.Split(ipv6AddrCidr, cidrSep)

	if len(splitted) != 2 {
		return false
	}

	prefix, suffix := splitted[0], splitted[1]
	number, consumed, ok := utils.Dtoi(suffix)

	if !ok || !Ipv6(prefix) || consumed > 3 {
		return false
	}

	return 0 <= number && number <= 128
}
