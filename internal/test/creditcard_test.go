package test

import (
	"testing"

	"github.com/eredotpkfr/golidators"
)

var CreditCardCases = [10]IntTestRecord{
	{1234567890, false},
	{1234567897, true},
	{5146713835430, false},
	{5146713835433, true},
	{5371087585041475, true},
	{5300025108592596, true},
	{5302025202593516, false},
	{5184214431476070, true},
	{5371087585041475, true},
	{5101365588025704, true},
}

func TestCreditCard(t *testing.T) {
	for _, record := range LuhnCases {
		if golidators.CreditCard(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
