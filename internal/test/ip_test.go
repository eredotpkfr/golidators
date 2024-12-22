package test

import (
	"testing"

	"github.com/eredotpkfr/golidators"
)

var (
	Ipv4Cases = [26]StrTestRecord{
		{"127.0.0.1", true},
		{"10.0.0.0", true},
		{"192.168.1.43", true},
		{"95.67.123.56", true},
		{"", false},
		{"123.5.77.88", true},
		{"12.12.12.12", true},
		{"255.255.255.255", true},
		{"abc.0.0.1", false},
		{"1278.0.0.1", false},
		{"127.0.0.abc", false},
		{"900.200.100.75", false},
		{"10.0.0.0:8080", false},
		{"foo/bar", false},
		{"127.0.0.1.", false},
		{".......", false},
		{"127.0.0.1....", false},
		{"...127.0.0.1...", false},
		{".127.0.0.1", false},
		{"foo", false},
		{"::1::2", false},
		{"::1", false},
		{"127.0.0.1.1.2.3", false},
		{".127.0.0.1.", false},
		{"clear.text", false},
		{"     ", false},
	}
	Ipv4CidrCases = [27]StrTestRecord{
		{"127.0.0.1/0", true},
		{"123.5.77.88/8", true},
		{"239.0.0.0/8", true},
		{"12.12.12.12/32", true},
		{"127.0.0.1/12", true},
		{"192.168.1.45/24", true},
		{"abc.0.0.1", false},
		{"1.1.1.1", false},
		{"1.1.1.1/-1", false},
		{"1.1.1.1/33", false},
		{"1.1.1.1/foo", false},
		{"127.0.b.a/12", false},
		{"127.0.0.2/44", false},
		{"a.b.c/123/", false},
		{".127.0.0.1/12", false},
		{"..127.0.0.1/12", false},
		{"127.0.0.1./12", false},
		{"127.0.0.1../12", false},
		{"", false},
		{"    ", false},
		{"foo", false},
		{"foo/bar", false},
		{"127.0.0.1/12/13/14", false},
		{"/", false},
		{"////", false},
		{"foo/bar/", false},
		{"/foo/bar", false},
	}
	Ipv6Cases = [20]StrTestRecord{
		{"::1", true},
		{"2002::", true},
		{"0:0:0:0:0:ffff:c0a8:12b", true},
		{"dead:beef:0:0:0:0:42:1", true},
		{"abcd:ef::42:1", true},
		{"0:0:0:0:0:ffff:1.2.3.4", true},
		{"2001:0db8:0000:0000:0000:ff00:0042:8329", true},
		{"::192.168.30.2", true},
		{"abc.0.0.1", false},
		{"2002::::::::::", false},
		{"abcd:1234::123::1", false},
		{"1:2:3:4:5:6:7:8:9", false},
		{"abcd::1ffff", false},
		{"", false},
		{"   ", false},
		{"foo", false},
		{"////", false},
		{":", false},
		{"::::", false},
		{"deag:beef:0:0:0:0:42:1", false},
	}
	Ipv6CidrCases = [26]StrTestRecord{
		{"::1/0", true},
		{"2002::/12", true},
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
		{"2002::::::::::/12", false},
		{"127.0.0.1/12/13/14", false},
		{"/", false},
		{"", false},
		{"   ", false},
		{"///", false},
		{"....", false},
		{"....", false},
		{"foo/bar/", false},
		{"/foo/bar", false},
	}
)

func TestIpv4(t *testing.T) {
	for _, record := range Ipv4Cases {
		if golidators.Ipv4(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv4Cidr(t *testing.T) {
	for _, record := range Ipv4CidrCases {
		if golidators.Ipv4Cidr(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv6(t *testing.T) {
	for _, record := range Ipv6Cases {
		if golidators.Ipv6(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestIpv6Cidr(t *testing.T) {
	for _, record := range Ipv6CidrCases {
		if golidators.Ipv6Cidr(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
