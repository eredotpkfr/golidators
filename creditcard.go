package golidators

// CreditCard validate credit card number with Luhn
func CreditCard(number int) bool {
	return Luhn(number)
}
