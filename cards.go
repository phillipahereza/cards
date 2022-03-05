package cards

import (
	"errors"
	"strconv"
	"strings"
)

const (
	amex       = "American Express"
	jcb        = "JCB"
	maestro    = "Maestro"
	visa       = "Visa"
	mastercard = "MasterCard"
	unknown    = "Unknown Card Scheme"
)

type cardDigits [4]int

// to returns the first n digits of the card number as an int
func (c *cardDigits) to(n int) int {
	return c[n-1]
}

func (c *cardDigits) assembleCardDigits(cardNumber string) error {
	for i := 0; i < 4; i++ {
		if i < len(cardNumber) {
			num, err := strconv.Atoi(cardNumber[:i+1])
			if err != nil {
				return errors.New("unknown credit card brand")
			}

			c[i] = num
		}
	}
	return nil
}

// IsCardNumberValid checks validity of a card number
func IsCardNumberValid(cardNumber string) bool {
	var sum, digit int
	var alternate bool
	var err error

	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	numberLen := len(cardNumber)

	if numberLen < 12 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		if digit, err = strconv.Atoi(string(cardNumber[i])); err != nil {
			return false
		}
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}

		alternate = !alternate
		sum += digit
	}

	return sum%10 == 0
}

// CardScheme returns the card scheme given a credit card number
func CardScheme(cardNumber string) string {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	digits := cardDigits{}
	err := digits.assembleCardDigits(cardNumber)
	if err != nil {
		return unknown
	}

	numberLength := len(cardNumber)

	if (strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37")) &&
		numberLength == 15 {
		return amex
	} else if strings.HasPrefix(cardNumber, "4") &&
		(numberLength == 13 || numberLength == 16 || numberLength == 19) {
		return visa
	} else if (digits.to(4) >= 3528 && digits.to(4) <= 3589) &&
		(numberLength >= 16 && numberLength <= 19) {
		return jcb
	} else if (strings.HasPrefix(cardNumber, "6") || strings.HasPrefix(cardNumber, "50") ||
		(digits.to(2) >= 56 && digits.to(2) <= 58)) && (numberLength >= 12 && numberLength <= 19) {
		return maestro
	} else if ((digits.to(4) >= 2221 && digits.to(4) <= 2720) || (digits.to(2) >= 51 && digits.to(2) <= 55)) &&
		numberLength == 16 {
		return mastercard
	}

	return unknown
}
