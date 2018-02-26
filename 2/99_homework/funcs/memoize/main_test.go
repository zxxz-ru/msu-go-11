package main

import (
	"testing"
)

func TestGetDigits(t *testing.T) {
	s := []struct {
		i   int
		n   int
		res string
	}{
		// ones
		{1, 1, "I"},
		{2, 1, "II"},
		{3, 1, "III"},
		{4, 1, "IV"},
		{5, 1, "V"},
		{6, 1, "VI"},
		{7, 1, "VII"},
		{8, 1, "VIII"},
		{9, 1, "IX"},
		// tens
		{1, 10, "X"},
		{2, 10, "XX"},
		{3, 10, "XXX"},
		{4, 10, "XL"},
		{5, 10, "L"},
		{6, 10, "LX"},
		{7, 10, "LXX"},
		{8, 10, "LXXX"},
		{9, 10, "XC"},
		// hundreds
		{1, 100, "C"},
		{2, 100, "CC"},
		{3, 100, "CCC"},
		{4, 100, "CD"},
		{5, 100, "D"},
		{6, 100, "DC"},
		{7, 100, "DCC"},
		{8, 100, "DCCC"},
		{9, 100, "CM"},
		// thousends
		{1, 1000, "M"},
		{2, 1000, "MM"},
		{3, 1000, "MMM"},
	}

	for _, c := range s {
		r := getDigits(c.i, c.n)
		if r != c.res {
			t.Errorf("expect %s, got %s", c.res, r)
		}
	}
}

func TestGetRoman(t *testing.T) {

	tt := []struct {
		i   int
		res string
	}{
		{1, "I"}, {2, "II"}, {3, "III"}, {4, "IV"}, {5, "V"}, {6, "VI"},
		{7, "VII"}, {8, "VIII"}, {9, "IX"}, {10, "X"}, {11, "XI"},
		{12, "XII"}, {13, "XIII"}, {14, "XIV"}, {15, "XV"}, {16, "XVI"},
		{17, "XVII"}, {18, "XVIII"}, {19, "XIX"}, {20, "XX"}, {25, "XXV"},
		{30, "XXX"}, {40, "XL"}, {50, "L"}, {60, "LX"}, {69, "LXIX"},
		{70, "LXX"}, {80, "LXXX"}, {90, "XC"}, {99, "XCIX"}, {100, "C"},
		{200, "CC"}, {300, "CCC"}, {400, "CD"}, {500, "D"}, {600, "DC"},
		{666, "DCLXVI"}, {700, "DCC"}, {800, "DCCC"}, {900, "CM"},
		{1000, "M"}, {1009, "MIX"}, {1444, "MCDXLIV"}, {1666, "MDCLXVI"},
		{1945, "MCMXLV"}, {1997, "MCMXCVII"}, {1999, "MCMXCIX"},
		{2000, "MM"}, {2008, "MMVIII"}, {2010, "MMX"}, {2012, "MMXII"},
		{2500, "MMD"}, {3000, "MMM"}, {3999, "MMMCMXCIX"}}
	for _, tc := range tt {

		r := getRoman(tc.i)
		if r != tc.res {
			t.Errorf("expect %s, got %s", tc.res, r)
		}
	}
}
