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
