package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(x []int) string {
	s := ""

	for i := range x {
		s += strconv.Itoa(x[i])
	}

	return s
}

func MergeSlices(x1 []float32, x2 []int32) []int {
	var x []int

	for i := range x1 {
		x = append(x, int(x1[i]))
	}

	for i := range x2 {
		x = append(x, int(x2[i]))
	}

	return x
}

func GetMapValuesSortedByKey(m map[int]string) []string {
	var keys []int

	for i, _ := range m {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	var s []string

	for i := range keys {
		s = append(s, m[keys[i]])
	}

	return s
}
