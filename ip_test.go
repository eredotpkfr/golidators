package golidators

import "testing"

var (
	Ipv4Cases = [10]StrTestRecord{
		StrTestRecord{"127.0.0.1", true},
		StrTestRecord{"10.0.0.0", true},
		StrTestRecord{"123.5.77.88", true},
		StrTestRecord{"12.12.12.12", true},
		StrTestRecord{"255.255.255.255", true},
		StrTestRecord{"abc.0.0.1", false},
		StrTestRecord{"1278.0.0.1", false},
		StrTestRecord{"127.0.0.abc", false},
		StrTestRecord{"900.200.100.75", false},
		StrTestRecord{"10.0.0.0:8080", false},
	}
	Ipv4CidrCases = [10]StrTestRecord{
		StrTestRecord{"127.0.0.1/0", true},
		StrTestRecord{"123.5.77.88/8", true},
		StrTestRecord{"12.12.12.12/32", true},
		StrTestRecord{"abc.0.0.1", false},
		StrTestRecord{"1.1.1.1", false},
		StrTestRecord{"1.1.1.1/-1", false},
		StrTestRecord{"1.1.1.1/33", false},
		StrTestRecord{"1.1.1.1/foo", false},
		StrTestRecord{"127.0.0.2/44", false},
		StrTestRecord{"foo/bar", false},
	}
	Ipv6Cases = [10]StrTestRecord{
		StrTestRecord{"::1", true},
		StrTestRecord{"0:0:0:0:0:ffff:c0a8:12b", true},
		StrTestRecord{"dead:beef:0:0:0:0:42:1", true},
		StrTestRecord{"abcd:ef::42:1", true},
		StrTestRecord{"0:0:0:0:0:ffff:1.2.3.4", true},
		StrTestRecord{"::192.168.30.2", true},
		StrTestRecord{"abc.0.0.1", false},
		StrTestRecord{"abcd:1234::123::1", false},
		StrTestRecord{"1:2:3:4:5:6:7:8:9", false},
		StrTestRecord{"abcd::1ffff", false},
	}
	Ipv6CidrCases = [15]StrTestRecord{
		StrTestRecord{"::1/0", true},
		StrTestRecord{"dead:beef:0:0:0:0:42:1/8", true},
		StrTestRecord{"abcd:ef::42:1/32", true},
		StrTestRecord{"0:0:0:0:0:ffff:1.2.3.4/64", true},
		StrTestRecord{"::192.168.30.2/128", true},
		StrTestRecord{"abc.0.0.1", false},
		StrTestRecord{"abcd:1234::123::1", false},
		StrTestRecord{"1:2:3:4:5:6:7:8:9", false},
		StrTestRecord{"abcd::1ffff", false},
		StrTestRecord{"1.1.1.1", false},
		StrTestRecord{"::1", false},
		StrTestRecord{"::1/129", false},
		StrTestRecord{"::1/-1", false},
		StrTestRecord{"::1/foo", false},
		StrTestRecord{"::1/144", false},
	}
)

func TestIpv4(t *testing.T) {
	for _, record := range Ipv4Cases {
		if Ipv4(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv4Cidr(t *testing.T) {
	for _, record := range Ipv4CidrCases {
		if Ipv4Cidr(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv6(t *testing.T) {
	for _, record := range Ipv6Cases {
		if Ipv6(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv6Cidr(t *testing.T) {
	for _, record := range Ipv6CidrCases {
		if Ipv6Cidr(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
