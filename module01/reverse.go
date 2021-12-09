package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reverseWord string
	for _, letter := range word {
		reverseWord = string(letter) + reverseWord
	}

	return reverseWord
}
