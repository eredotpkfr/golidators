package test

import (
	"testing"

	"github.com/eredotpkfr/golidators"
)

var PortCases = [5]IntTestRecord{
	{12, true},
	{80, true},
	{443, true},
	{99999, false},
	{65536, false},
}

func TestPort(t *testing.T) {
	for _, record := range PortCases {
		if golidators.Port(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
