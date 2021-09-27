package golidators

import "testing"

var PortCases = [5]IntTestRecord{
	IntTestRecord{12, true},
	IntTestRecord{80, true},
	IntTestRecord{443, true},
	IntTestRecord{99999, false},
	IntTestRecord{65536, false},
}

func TestPort(t *testing.T) {
	for _, record := range PortCases {
		if Port(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
