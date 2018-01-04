package bubble

import (
	"fmt"
)

func BubbleSort(lst ...interface{}) {
	bound := len(lst) - 1
	for {
		t := 0
		for j := 0; j < bound; j++ {
			if fmt.Sprintf("%v", lst[j]) > fmt.Sprintf("%v", lst[j+1]) {
				lst[j], lst[j+1] = lst[j+1], lst[j]
				t = j
			}
		}
		if t == 0 {
			break
		}
		bound = t
	}
}
