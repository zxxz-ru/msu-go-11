package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// vars
var(
    romans map[int]string = getMap()
)

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

// return arg i multiply on arg n as string in Roman digits
// arg i is number of ones or tens, or hundreds
// n flag to signal what is it must be [1, 10, 100, 1000]

    func GetRoman (i int) interface{} {
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
