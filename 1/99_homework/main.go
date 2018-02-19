package main

import (
	"fmt"
	"sort"
)

const (
	const_i = int(1)
	const_f = float32(1.1)
)

// always return  1
func ReturnInt() int {
	return const_i
}

// always return 1.1
func ReturnFloat() float32 {
	return const_f
}

// In func belowe no one can change value of local var, thus no matter what
// funcs will always return expected value. So we created kinda constant
///slice and array. Cool!

//  always return [3]int{1, 3, 4}
func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}

}

// always return []int{1, 2, 3}
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

// return string of concatinated slice of int
func IntSliceToString(slc []int) string {
	s := ""
	for _, i := range slc {
		// simple string + int doesn't work
		s = s + fmt.Sprintf("%d", i)
	}
	return s
}

// get slice of float32 and slice of int32 merge them and return slice
// of int
func MergeSlices(sf []float32, si []int32) []int {
	size := len(sf) + len(si)
	res := make([]int, 0, size)
	for _, f := range sf {
		res = append(res, int(f))
	}
	for _, i := range si {
		res = append(res, int(i))
	}
	return res
}

// return string slice of values sorted by map int key
// use sort pkg
func GetMapValuesSortedByKey(m map[int]string) []string {
	size := len(m)
	res := make([]string, 0, size)
	keys := make([]int, 0, size)
	for key := range m {
		keys = append(keys, key)
	}

	var sort_keys sort.IntSlice
	sort_keys = keys
	sort_keys.Sort()

	for _, key := range sort_keys {
		res = append(res, m[key])
	}
	return res
}

func main() {
	m := map[int]string{123: "a", 54: "b", 1: "c"}
	str := GetMapValuesSortedByKey(m)
	fmt.Printf("%v\n", str)
}
