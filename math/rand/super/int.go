package super

import "math/rand"

type __SuperRandInt__ struct {
	weight int // 权重
	min    int // 最小值
	max    int // 最大值
}
type __SuperRandIntSlice__ []__SuperRandInt__

func (slice __SuperRandIntSlice__) __rand__() int {
	sliceLen := len(slice)
	weightSum := 0
	for i := 0; i < sliceLen; i++ {
		weightSum += slice[i].weight
	}

	randVal := rand.Intn(weightSum)
	result := 0
	for i := 0; i < sliceLen; i++ {
		if randVal <= slice[i].weight {
			result = __intNM__(slice[i].min, slice[i].max)
			break
		}
		randVal -= slice[i].weight
	}
	return result
}

func __intNM__(n int, m int) int {
	return n + rand.Intn(m-n)
}

func __init__() __SuperRandIntSlice__ {
	std := make(__SuperRandIntSlice__, 0)
	return std
}
func (std __SuperRandIntSlice__) __add__(weight int, min int, max int) __SuperRandIntSlice__ {
	std = append(std, __SuperRandInt__{weight, min, max})
	return std
}
