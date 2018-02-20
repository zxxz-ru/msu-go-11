package main

import (
	"testing"
)

func TestGet_digits(t *testing.T) {
	s := []struct {
		i, n int
		res  string
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
		// {10,""},
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
		r, err := get_digits(c.i, c.n)
		if err != nil {
			t.Errorf("func responded with error: %v", err)
		}
		if r != c.res {
			t.Errorf("expect %s, got %s", c.res, r)
		}
	}
}
