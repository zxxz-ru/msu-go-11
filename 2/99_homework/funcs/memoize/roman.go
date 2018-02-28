package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	romans map[int]string = getMap()
)

// always will return constant map of int keys to Roman digit.
func getMap() map[int]string {
	return map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
}

// Convert integer less than 4000 to Roman digits.
// does not check for negative and zero.
func getRoman(i int) interface{} {
	str := strconv.Itoa(i)
	n := 1
	length := len(str)
	slc := make([]string, length, length)
	for i := length - 1; i >= 0; i-- {
		d, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		s := getDigits(d, n)
		n *= 10
		slc[i] = s
	}
	resultStr := strings.Join(slc, "")
	var itr interface{} = resultStr

	return itr
}

// return roman representation of arabic digit
// first argument is a digit to convert, second
// argument must be either 1, 10, 100, 1000.
// For example pass 5, 10 func will convert 50.
// Pass 5, 1 func will convert 5.
func getDigits(i, n int) (res string) {
	switch {
	case i < 4:
		for i > 0 {
			res += romans[n]
			i = i - 1
		}
	case i == 4:
		res = romans[(4 * n)]
	case i == 5:
		res = romans[(5 * n)]
	case 5 < i && i < 9:
		res += romans[(5 * n)]
		i = i - 5
		for i > 0 {
			res += romans[n]
			i = i - 1
		}
	case i == 9:
		res += romans[(9 * n)]
	}
	return
}
