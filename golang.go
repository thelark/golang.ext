package main

import (
	"fmt"
	"thelark.cn/golang.ext/sort"
	s "sort"
)

/**
 * TODO: golang 扩展包
 */

func main() {
	a := sort.IntSlice{1, 3, 5, 7, 2}
	b := []float64{1.1, 2.3, 5.3, 3.4}
	c := []int{1, 3, 5, 4, 2}
	fmt.Println(s.IsSorted(a)) //false
	if !s.IsSorted(a) {
		a.Sort()
	}

	if !s.Float64sAreSorted(b) {
		s.Float64s(b)
	}
	if !s.IntsAreSorted(c) {
		s.Ints(c)
	}
	fmt.Println(a)//[1 2 3 5 7]
	fmt.Println(b)//[1.1 2.3 3.4 5.3]
	fmt.Println(c)// [1 2 3 4 5]
}