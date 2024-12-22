package golidators

func luhnChecksum(number int) int {
	var sum int = 0

	for i := 0; number > 0; i++ {
		digit := number % 10

		if i%2 == 0 {
			digit = digit * 2

			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}

		sum += digit
		number = number / 10
	}

	return sum % 10
}

// Luhn checks if the given integer is valid according to the Luhn algorithm
func Luhn(number int) bool {
	return (number%10+luhnChecksum(number/10))%10 == 0
}

// LuhnCheckDigit calculates the check digit for a given number using the Luhn algorithm
func LuhnCheckDigit(number int) int {
	return (10 - luhnChecksum(number)) % 10
}
