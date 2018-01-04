package rand

func NewRandIntSlice() __SuperRandIntSlice__ {
	return __init__()
}
func (std __SuperRandIntSlice__) Add(weight int, min int, max int) __SuperRandIntSlice__ {
	return std.__add__(weight, min, max)
}
func (std __SuperRandIntSlice__) Rand() int {
	return std.__rand__()
}
func (std __SuperRandIntSlice__) IntNM(n, m int) int {
	return std.IntNM(n, m)
}
