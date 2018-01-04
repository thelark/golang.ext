package sort

import (
	"fmt"
	"sort"
)

// ASC 正序
// DESC 倒序
type sortRule uint

const (
	ASC  sortRule = iota
	DESC
)

func NewSlice(datas ...interface{}) interfaceSlice {
	return interfaceSlice(datas)
}

var defaultSortRule = ASC
//定义interface{},并实现sort.Interface接口的三个方法
type interfaceSlice []interface{}

func (c interfaceSlice) Len() int {
	return len(c)
}
func (c interfaceSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c interfaceSlice) Less(i, j int) bool {
	switch defaultSortRule {
	case ASC:
		return fmt.Sprintf("%v", c[i]) < fmt.Sprintf("%v", c[j])
	case DESC:
		return fmt.Sprintf("%v", c[i]) > fmt.Sprintf("%v", c[j])
	default:
		return fmt.Sprintf("%v", c[i]) < fmt.Sprintf("%v", c[j])
	}
}

func (c interfaceSlice) Sort(rule ...sortRule) error {
	if len(rule) > 1 {
		return fmt.Errorf("传入了过多的排序规则...")
	}
	if len(rule) == 1 {
		defaultSortRule = rule[0]
	}
	sort.Sort(c)
	return nil
}
