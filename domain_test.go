package golidators

import "testing"

var DomainCases = [35]StrTestRecord{
	StrTestRecord{"example.com", true},
	StrTestRecord{"EXAMPLE.com", true},
	StrTestRecord{"xn----gtbspbbmkef.xn--p1ai", true},
	StrTestRecord{"underscore_subdomain.example.com", true},
	StrTestRecord{"something.versicherung", true},
	StrTestRecord{"someThing.versicherung", true},
	StrTestRecord{"11.com", true},
	StrTestRecord{"11.COM", true},
	StrTestRecord{"3.cn", true},
	StrTestRecord{"a.cn", true},
	StrTestRecord{"sub1.sub2.sample.co.uk", true},
	StrTestRecord{"somerandomexample.xn--fiqs8s", true},
	StrTestRecord{"kräuter.com", true},
	StrTestRecord{"über.com", true},
	StrTestRecord{"xn--eckwd4c7c.xn--zckzah", true},
	StrTestRecord{"XN--ECKWD4C7C.XN--ZCKZAH", true},
	StrTestRecord{"www.google.com", true},
	StrTestRecord{"google.com", true},
	StrTestRecord{"example.com/", false},
	StrTestRecord{"example.com:4444", false},
	StrTestRecord{"example.-com", false},
	StrTestRecord{"example.", false},
	StrTestRecord{"-example.com", false},
	StrTestRecord{"example-.com", false},
	StrTestRecord{"_example.com", false},
	StrTestRecord{"example_.com", false},
	StrTestRecord{"EXAMPLE_.COM", false},
	StrTestRecord{"example", false},
	StrTestRecord{"a......b.com", false},
	StrTestRecord{"A......B.COM", false},
	StrTestRecord{"a.123", false},
	StrTestRecord{"123.123", false},
	StrTestRecord{"123.123.123", false},
	StrTestRecord{"123.123.123.123", false},
	StrTestRecord{"example.com:4444", false},
}

func TestDomain(t *testing.T) {
	for _, record := range DomainCases {
		if Domain(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
