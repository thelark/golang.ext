package bubble

import (
	"fmt"
	"time"
)

func BubbleSort(lst []int) {
	start := time.Now().UnixNano()
	bound := len(lst) - 1
	for {
		t := 0
		for j := 0; j < bound; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j]
				t = j
			}
		}
		if t == 0 {
			break
		}
		bound = t
	}
	fmt.Printf("耗时: %v纳秒\n", time.Now().UnixNano()-start)
	fmt.Printf("耗时: %v毫秒\n", (time.Now().UnixNano()-start)/1e6) //将纳秒转换为毫秒
	fmt.Printf("耗时: %v秒\n", (time.Now().UnixNano()-start)/1e9)  //将纳秒转换为秒
}
