package golidators

import "testing"

var (
	Ipv4Cases = [10]StrTestRecord{
		{"127.0.0.1", true},
		{"10.0.0.0", true},
		{"123.5.77.88", true},
		{"12.12.12.12", true},
		{"255.255.255.255", true},
		{"abc.0.0.1", false},
		{"1278.0.0.1", false},
		{"127.0.0.abc", false},
		{"900.200.100.75", false},
		{"10.0.0.0:8080", false},
	}
	Ipv4CidrCases = [10]StrTestRecord{
		{"127.0.0.1/0", true},
		{"123.5.77.88/8", true},
		{"12.12.12.12/32", true},
		{"abc.0.0.1", false},
		{"1.1.1.1", false},
		{"1.1.1.1/-1", false},
		{"1.1.1.1/33", false},
		{"1.1.1.1/foo", false},
		{"127.0.0.2/44", false},
		{"foo/bar", false},
	}
	Ipv6Cases = [10]StrTestRecord{
		{"::1", true},
		{"0:0:0:0:0:ffff:c0a8:12b", true},
		{"dead:beef:0:0:0:0:42:1", true},
		{"abcd:ef::42:1", true},
		{"0:0:0:0:0:ffff:1.2.3.4", true},
		{"::192.168.30.2", true},
		{"abc.0.0.1", false},
		{"abcd:1234::123::1", false},
		{"1:2:3:4:5:6:7:8:9", false},
		{"abcd::1ffff", false},
	}
	Ipv6CidrCases = [15]StrTestRecord{
		{"::1/0", true},
		{"dead:beef:0:0:0:0:42:1/8", true},
		{"abcd:ef::42:1/32", true},
		{"0:0:0:0:0:ffff:1.2.3.4/64", true},
		{"::192.168.30.2/128", true},
		{"abc.0.0.1", false},
		{"abcd:1234::123::1", false},
		{"1:2:3:4:5:6:7:8:9", false},
		{"abcd::1ffff", false},
		{"1.1.1.1", false},
		{"::1", false},
		{"::1/129", false},
		{"::1/-1", false},
		{"::1/foo", false},
		{"::1/144", false},
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
