package main

import (
	"fmt"
)

// https://ru.wikipedia.org/wiki/Алгоритм_нахождения_корня_n-ной_степени)
// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	var a, a_k float64
	f := func() (flt float64) {
		flt = (0.5 * (a + (x / a)))
		return
	}

	a = x / 2
	a_k = f()
    for i := 0; i < 7; i++ {
		a = a_k
		a_k = f()
	}
	return a_k
}

func main() {
	fmt.Println(Sqrt(25))
}
