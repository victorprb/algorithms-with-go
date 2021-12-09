package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	values := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			values = append(values, "Fizz Buzz")
		case i%5 == 0:
			values = append(values, "Buzz")
		case i%3 == 0:
			values = append(values, "Fizz")
		default:
			values = append(values, fmt.Sprint(i))
		}
	}
	fmt.Println(strings.Join(values, ", "))
}
