package module01

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14

const charset = "0123456789ABCDEF"

func BaseToDec(value string, base int) int {
	result := 0
	value = Reverse(value)

	for i, v := range value {
		convertedInt := convertValueToInt(string(v))
		result = result + (convertedInt * int(math.Pow(float64(base), float64(i))))
	}

	return result
}

func convertValueToInt(input string) int {
	for i, v := range charset {
		if string(v) == input {
			return i
		}
	}
	return 0
}
