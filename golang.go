package main

import (
	"fmt"
	s "sort"
	"thelark.cn/golang.ext/sort"
	"thelark.cn/golang.ext/sort/bubble"
)

/**
 * TODO: golang 扩展包
 */

func main() {
	a := sort.NewSlice(1, 3, 5, 7, 2)
	b := []float64{1.1, 2.3, 5.3, 3.4}
	c := []int{1, 3, 5, 4, 2}

	xxb := make([]interface{}, len(b))
	for i := range b {
		xxb[i] = b[i]
	}

	xxc := make([]interface{}, len(c))
	for i := range c {
		xxc[i] = c[i]
	}
	fmt.Println(xxb)
	bubble.BubbleSort(xxb...)
	fmt.Println(xxb)
	return
	xb := sort.NewSlice(xxb...)
	xc := sort.NewSlice(xxc...)

	fmt.Println(s.IsSorted(a)) //false
	if !s.IsSorted(a) {
		a.Sort()
	}
	fmt.Println(s.IsSorted(a)) //false
	if !s.Float64sAreSorted(b) {
		s.Float64s(b)
	}
	if !s.IntsAreSorted(c) {
		s.Ints(c)
	}
	fmt.Println("------------")
	fmt.Println(xb)
	fmt.Println(xc)
	xb.Sort()
	xc.Sort()
	fmt.Println(xb)
	fmt.Println(xc)
	fmt.Println("------------")
	fmt.Println(a) //[1 2 3 5 7]
	fmt.Println(b) //[1.1 2.3 3.4 5.3]
	fmt.Println(c) // [1 2 3 4 5]
}
