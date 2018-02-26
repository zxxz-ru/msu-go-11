package main

import (
	"fmt"
    "os"
    "strconv"
    "strings"
)
// TODO write test for getRomams
// TODO remake getRomans as a for loop and string conversion
// vars
var romans map[int]string

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

var cache map[int]int = make(map[int]int)

// ...int slice of keys?
type memoizeFunction func(int, ...int) interface{}

// return arg i multiply on arg n as string in Roman digits
// arg i is number of ones or tens, or hundreds
// n flag to signal what is it must be [1, 10, 100, 1000]
func getRoman(i int) ( string) {
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
    return strings.Join(slc,"")
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

// func getRoman(i int) (res string) {
// 	switch {
// 	case i >= 1000:
// 		th := int(i / 1000)
// 		i = i - (1000 * th)
// 		res += get_digits(th, 1000)
// 		//fmt.Println(res)
// 		res += getRoman(i)
// 	case i >= 100:
// 		hrd := int(i / 100)
// 		i = i - (100 * hrd)
// 		res += get_digits(hrd, 100)
// 		//fmt.Println(res)
// 		res += getRoman(i)
// 
// 	case i >= 10:
// 		tns := int(i / 10)
// 		i = i - (10 * tns)
// 		res += get_digits(tns, 10)
// 		//fmt.Println(res)
// 		res += getRoman(i)
// 	case i < 10:
// 		res += get_digits(i, 1)
// 		//fmt.Println(res)
// 	}
// 	return
// }

// TODO реализовать
// var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	return function
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	romans = getMap()
}

func main() {
	// 	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	/*
		for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
			14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
			90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
			1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
			2012, 2500, 3000, 3999} {
			fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	    }
	*/

	fmt.Println(getRoman(1975))
}
