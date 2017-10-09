package sort

import (
	"sort"
	"fmt"
)

// ASC 正序
// DESC 倒序
type sortRule uint

const (
	ASC  sortRule = iota
	DESC
)

//定义interface{},并实现sort.Interface接口的三个方法
type IntSlice []interface{}

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	switch defaultSortRule {
	case ASC:
		return fmt.Sprintf("%v", c[i]) < fmt.Sprintf("%v", c[j])
	case DESC:
		return fmt.Sprintf("%v", c[i]) > fmt.Sprintf("%v", c[j])
	default:
		return fmt.Sprintf("%v", c[i]) < fmt.Sprintf("%v", c[j])
	}
}

var defaultSortRule = ASC

func (c IntSlice) Sort(rule ...sortRule) error {
	if len(rule) > 1 {
		return fmt.Errorf("传入了过多的排序规则...")
	}
	if len(rule) == 1 {
		defaultSortRule = rule[0]
	}
	sort.Sort(c)
	return nil
}
