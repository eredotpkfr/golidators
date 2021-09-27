package golidators

import "testing"

var MacCases = [7]StrTestRecord{
	StrTestRecord{"01:23:45:67:ab:CD", true},
	StrTestRecord{"01:23:45:67:AB:CD", true},
	StrTestRecord{"00:00:00:00:00", false},
	StrTestRecord{"01:23:45:67:89:", false},
	StrTestRecord{"01:23:45:67:89:gh", false},
	StrTestRecord{"123:23:45:67:89:00", false},
	StrTestRecord{"01:23:45:67:89:GH", false},
}

func TestMac(t *testing.T) {
	for _, record := range MacCases {
		if Mac(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
