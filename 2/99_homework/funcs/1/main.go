package main
import (
    "sort"
    "fmt"
}
const romans = map[int] string{
    1 : "I",
    4 : "IV",
    5 : "V",
    9 : "IX",
    10 : "X",
    40 : "XL",
    50 : "L",
    90 : "XC",
    100 : "C",
    400 : "CD",
    500 : "D",
    900 : "CM",
    1000 : "M",
}


var cache map[int] int = make(map[int] int, 0,4)
// ...int slice of keys?
type memoizeFunction func(int, ...int) interface{}

// return roman digit for integers less than 10 as string
func get_ones(i int) (res string) {
    swith {
    case i < 4:
        for i; i>0; i--{
res += romans[1]
        }
    case i == 4:
        res = romans[4]
    case i == 5:
        res = romans[5]
    case  5 < i <9:
        res += romans[5]
        for i = i-5; i > 0; i-- {
res += romans[1]
        }
    case i == 9:
        res += romans[9]
    }

    return
}
func get_tens(i int) (res string) {
    

    return
}

func get_hundreds(i int) (res string) {
    

    return
}
func get_ones(i int) (res string) {
    

    return
}

// TODO реализовать
// var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	return function
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
}

func main() {
// 	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
