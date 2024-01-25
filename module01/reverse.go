package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"

// THIS IS CORRECT AND IT DEALS WITH STRANGE CHARS
func ReverseEasy(word string) string {
	reversedWord := ""
	for _, c := range word {
		reversedWord = string(c) + reversedWord
	}
	return reversedWord
}

// THIS IS INCORRECT IF YOU USE STRANGE CHARS
func Reverse(word string) string {
	if len(word) == 0 {
		return ""
	}
	return Reverse(word[1:]) + string(word[0])
}
