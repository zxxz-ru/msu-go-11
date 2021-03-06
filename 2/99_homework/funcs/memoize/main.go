package main

import (
	"fmt"
)

type memoizeFunction func(int) interface{}

var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

func memoize(function memoizeFunction) memoizeFunction {
	var cache map[int]interface{} = make(map[int]interface{})
	return func(i int) (res interface{}) {
		if r, ok := cache[i]; ok {
			res = r
		} else {
			r := function(i)
			cache[i] = r
			res = r
		}
		return
	}
}

// func fib
func fib(i int) (itf interface{}) {
	sl := []int{}
	x, y := 0, 1
	func() {
		for x <= i {
			sl = append(sl, x)
			x, y = y, x+y
		}
	}()
	itf = sl
	return
}

func init() {
	romanForDecimal = memoize(getRoman)
	fibonacci = memoize(fib)
}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).([]int))

	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 111,
		/*
			    14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
					90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
					1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
					2012, 2500, 3000, 3999
		*/

	} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
