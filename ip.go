package golidators

import utils "github.com/eredotpkfr/golidators/internal/utilities"

const (
	// IP CIDR seperator
	cidr_sep = '/'

	// 0xFF -> 255
	ipv4_max_part = 0xFF
	ipv4_length   = 4

	// 0xFFFF -> 65535
	ipv6_max_part = 0xFFFF
	ipv6_length   = 16
)

func Ipv4(ipv4_addr string) bool {
	for index := 0; index < ipv4_length; index++ {

		if len(ipv4_addr) == 0 {
			return false
		}

		if index > 0 {
			if ipv4_addr[0] != '.' {
				return false
			}

			ipv4_addr = ipv4_addr[1:]
		}

		number, consumed, ok := utils.Dtoi(ipv4_addr)

		if !ok || number > ipv4_max_part {
			return false
		}

		if consumed > 1 && ipv4_addr[0] == '0' {
			return false
		}

		ipv4_addr = ipv4_addr[consumed:]

	}

	if len(ipv4_addr) != 0 {
		return false
	}

	return true
}

func Ipv4Cidr(ipv4_addr_cidr string) bool {
	prefix, suffix, ok := utils.SplitOnce(ipv4_addr_cidr, cidr_sep)

	if !ok {
		return false
	}

	number, consumed, ok := utils.Dtoi(suffix)

	if !ok || !Ipv4(prefix) || consumed > 2 {
		return false
	}

	return 0 <= number && number <= 32
}

func Ipv6(ipv6_addr string) bool {
	ellipsis := -1

	if len(ipv6_addr) >= 2 && ipv6_addr[0] == ':' && ipv6_addr[1] == ':' {
		ellipsis = 0
		ipv6_addr = ipv6_addr[2:]

		if len(ipv6_addr) == 0 {
			return true
		}
	}

	index := 0
	for index < ipv6_length {
		number, consumed, ok := utils.Xtoi(ipv6_addr)

		if !ok || number > ipv6_max_part {
			return false
		}

		if consumed < len(ipv6_addr) && ipv6_addr[consumed] == '.' {
			if ellipsis < 0 && index != ipv6_length-ipv4_length {
				return false
			}
			if index+ipv4_length > ipv6_length {
				return false
			}
			if !Ipv4(ipv6_addr) {
				return false
			}

			ipv6_addr = ""
			index += ipv4_length
			break
		}

		index += 2

		ipv6_addr = ipv6_addr[consumed:]
		if len(ipv6_addr) == 0 {
			break
		}

		if ipv6_addr[0] != ':' || len(ipv6_addr) == 1 {
			return false
		}

		ipv6_addr = ipv6_addr[1:]

		if ipv6_addr[0] == ':' {
			if ellipsis >= 0 {
				return false
			}

			ellipsis = index
			ipv6_addr = ipv6_addr[1:]

			if len(ipv6_addr) == 0 {
				break
			}

		}
	}

	if len(ipv6_addr) != 0 {
		return false
	}

	if index < ipv6_length {
		if ellipsis < 0 {
			return false
		}
	} else if ellipsis >= 0 {
		return false
	}

	return true
}

func Ipv6Cidr(ipv6_addr_cidr string) bool {
	prefix, suffix, ok := utils.SplitOnce(ipv6_addr_cidr, cidr_sep)

	if !ok {
		return false
	}

	number, consumed, ok := utils.Dtoi(suffix)

	if !ok || !Ipv6(prefix) || consumed > 3 {
		return false
	}

	return 0 <= number && number <= 128
}
