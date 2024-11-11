package propertybasedtests

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		arabicNumber int
		Want         string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{13, "XIII"},
		{14, "XIV"},
		{19, "XIX"},
		{24, "XXIV"},
		{29, "XXIX"},
		{31, "XXXI"},
		{38, "XXXVIII"},
		{40, "XL"},
		{48, "XLVIII"},
		{49, "XLIX"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{399, "CCCXCIX"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1001, "MI"},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := ConvertToRoman(test.arabicNumber)
			if got != test.Want {
				t.Errorf("\ngot  %q \nwant %q", got, test.Want)
			}
		})
	}

}
