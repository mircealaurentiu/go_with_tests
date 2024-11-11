package propertybasedtests

import "strings"

func ConvertToRoman(arabicNumber int) string {

	var result strings.Builder

	for arabicNumber > 0 {
		switch { // switch on true

		case arabicNumber > 9:
			result.WriteString("X")
			arabicNumber -= 10

		case arabicNumber > 8:
			result.WriteString("IX")
			arabicNumber -= 9

		case arabicNumber > 4:
			result.WriteString("V")
			arabicNumber -= 5

		case arabicNumber > 3:
			result.WriteString("IV")
			arabicNumber -= 4

		default:
			result.WriteString("I")
			arabicNumber--
		}
	}
	return result.String()

}
