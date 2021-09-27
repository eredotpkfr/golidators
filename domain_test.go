package golidators

import "testing"

var DomainCases = [35]StrTestRecord{
	{"example.com", true},
	{"EXAMPLE.com", true},
	{"xn----gtbspbbmkef.xn--p1ai", true},
	{"underscore_subdomain.example.com", true},
	{"something.versicherung", true},
	{"someThing.versicherung", true},
	{"11.com", true},
	{"11.COM", true},
	{"3.cn", true},
	{"a.cn", true},
	{"sub1.sub2.sample.co.uk", true},
	{"somerandomexample.xn--fiqs8s", true},
	{"kräuter.com", true},
	{"über.com", true},
	{"xn--eckwd4c7c.xn--zckzah", true},
	{"XN--ECKWD4C7C.XN--ZCKZAH", true},
	{"www.google.com", true},
	{"google.com", true},
	{"example.com/", false},
	{"example.com:4444", false},
	{"example.-com", false},
	{"example.", false},
	{"-example.com", false},
	{"example-.com", false},
	{"_example.com", false},
	{"example_.com", false},
	{"EXAMPLE_.COM", false},
	{"example", false},
	{"a......b.com", false},
	{"A......B.COM", false},
	{"a.123", false},
	{"123.123", false},
	{"123.123.123", false},
	{"123.123.123.123", false},
	{"example.com:4444", false},
}

func TestDomain(t *testing.T) {
	for _, record := range DomainCases {
		if Domain(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
