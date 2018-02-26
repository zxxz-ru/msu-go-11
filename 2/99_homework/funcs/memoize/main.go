package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type memoizeFunction func(int) string

// vars
var romans map[int]string

// var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

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
func getRoman(i int) string {
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
	return strings.Join(slc, "")
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

//TODO Write memoization function

// ...int slice of keys?

func memoize(function memoizeFunction) memoizeFunction {
	var cache map[int]string = make(map[int]string)
	f := func(i int) (res string) {
		if r, ok := cache[i]; ok {
			res = r
		} else {
			r := function(i)
			cache[i] = r
			res = r
		}
		return
	}
	return f
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	romans = getMap()
	romanForDecimal = memoize(getRoman)
}

func main() {
	// 	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		/*
		    14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
				90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
				1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
				2012, 2500, 3000, 3999
		*/
	} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x))
	}
}
