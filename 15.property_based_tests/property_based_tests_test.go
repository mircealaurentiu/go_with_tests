package propertybasedtests

import (
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
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

func TestConvertToRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("\ngot  %q \nwant %q", got, test.Roman)
			}
		})
	}

}

func TestConvertToArabic(t *testing.T) {

	for _, test := range cases[:4] {
		t.Run("", func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("\ngot  %d \nwant %d", got, test.Arabic)
			}
		})
	}

}

func TestPropertiesOfConversion(t *testing.T) {

	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	err := quick.Check(assertion, nil)
	if err != nil {
		t.Error("fail check", err)
	}
}
