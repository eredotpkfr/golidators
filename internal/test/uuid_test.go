package test

import (
	"github.com/eredotpkfr/golidators"
	"testing"
)

var UUIDCases = [7]StrTestRecord{
	{"2bc1c94f-0deb-43e9-92a1-4775189ec9f8", true},
	{"d22877be-b0cf-4d81-99bf-242afba289a6", true},
	{"D22877BE-B0CF-4D81-99BF-242AFBA289A6", true},
	{"2bc1c94f-deb-43e9-92a1-4775189ec9f8", false},
	{"2bc1c94f-0deb-43e9-92a1-4775189ec9f", false},
	{"gbc1c94f-0deb-43e9-92a1-4775189ec9f8", false},
	{"2bc1c94f 0deb-43e9-92a1-4775189ec9f8", false},
}

func TestUuid(t *testing.T) {
	for _, record := range UUIDCases {
		if golidators.Uuid(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
