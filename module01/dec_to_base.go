package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"

const charSet = "0123456789ABCDEF"

func DecToBase(dec, base int) string {
	if dec == 0 {
		return ""
	}

	reminder := dec % base

	dec = dec / base
	return DecToBase(dec, base) + string(charSet[reminder])

}
