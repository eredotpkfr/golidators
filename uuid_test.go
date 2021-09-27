package golidators

import "testing"

var UuidCases = [7]StrTestRecord{
	StrTestRecord{"2bc1c94f-0deb-43e9-92a1-4775189ec9f8", true},
	StrTestRecord{"d22877be-b0cf-4d81-99bf-242afba289a6", true},
	StrTestRecord{"D22877BE-B0CF-4D81-99BF-242AFBA289A6", true},
	StrTestRecord{"2bc1c94f-deb-43e9-92a1-4775189ec9f8", false},
	StrTestRecord{"2bc1c94f-0deb-43e9-92a1-4775189ec9f", false},
	StrTestRecord{"gbc1c94f-0deb-43e9-92a1-4775189ec9f8", false},
	StrTestRecord{"2bc1c94f 0deb-43e9-92a1-4775189ec9f8", false},
}

func TestUuid(t *testing.T) {
	for _, record := range UuidCases {
		if Uuid(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
