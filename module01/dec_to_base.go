package module01

import (
	"fmt"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	baseDecimal := fmt.Sprint(dec % base)

	switch baseDecimal {
	case "10":
		baseDecimal = "A"
	case "11":
		baseDecimal = "B"
	case "12":
		baseDecimal = "C"
	case "13":
		baseDecimal = "D"
	case "14":
		baseDecimal = "E"
	case "15":
		baseDecimal = "F"
	}

	dec = dec / base

	if dec <= 0 {
		return baseDecimal
	}

	return DecToBase(dec, base) + baseDecimal
}
