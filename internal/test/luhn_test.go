package test

import (
	"testing"

	"github.com/eredotpkfr/golidators"
)

var LuhnCases = [7]IntTestRecord{
	{1234567890, false},
	{1234567897, true},
	{5146713835430, false},
	{5146713835433, true},
	{5371087585041475, true},
	{5300025108592596, true},
	{5302025202593516, false},
}

var LuhnCheckDigitCases = [7]IntTestRecordWithIntReturn{
	{1234567890, 3},
	{1234567897, 8},
	{5146713835430, 2},
	{5146713835433, 6},
	{5371087585041475, 1},
	{5300025108592596, 2},
	{5302025202593516, 6},
}

func TestLuhn(t *testing.T) {
	for _, record := range LuhnCases {
		if golidators.Luhn(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestLuhnCheckDigit(t *testing.T) {
	for _, record := range LuhnCheckDigitCases {
		if golidators.LuhnCheckDigit(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
