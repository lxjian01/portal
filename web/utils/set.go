package utils

import (
	"sort"
)

type Set map[int]bool

// 新建集合对象
// 可以传入初始元素
func New(items ...int) Set {
	s := make(Set, len(items))
	s.Add(items...)
	return s
}

// 创建副本
func (s Set) Duplicate() Set {
	r := make(map[int]bool, len(s))
	for e := range s {
		r[e] = true
	}
	return r
}

// 添加元素
func (s Set) Add(items ...int) {
	for _, v := range items {
		s[v] = true
	}
}

// 删除元素
func (s Set) Remove(items ...int) {
	for _, v := range items {
		delete(s, v)
	}
}

// 判断元素是否存在
func (s Set) Has(items ...int) bool {
	for _, v := range items {
		if _, ok := s[v]; !ok {
			return false
		}
	}
	return true
}

// 统计元素个数
func (s Set) Count() int {
	return len(s)
}

// 清空集合
func (s Set) Clear() {
	s = map[int]bool{}
}

// 空集合判断
func (s Set) Empty() bool {
	return len(s) == 0
}

// 获取元素列表（无序）
func (s Set) List() []int {
	list := make([]int, 0, len(s))
	for item := range s {
		list = append(list, item)
	}
	return list
}

// 获取元素列表（有序）
func (s Set) SortedList() []int {
	list := s.List()
	sort.Ints(list)
	return list
}

// 并集
// 获取 s 与参数的并集，结果存入 s
func (s Set) Union(sets ...Set) {
	for _, set := range sets {
		for e := range set {
			s[e] = true
		}
	}
}

// 并集（函数）
// 获取所有参数的并集，并返回
func Union(sets ...Set) Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取并集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range set {
			r[e] = true
		}
	}
	return r
}

// 差集
// 获取 s 与所有参数的差集，结果存入 s
func (s Set) Minus(sets ...Set) {
	for _, set := range sets {
		for e := range set {
			delete(s, e)
		}
	}
}

// 差集（函数）
// 获取第 1 个参数与其它参数的差集，并返回
func Minus(sets ...Set) Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取差集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range set {
			delete(r, e)
		}
	}
	return r
}

// 交集
// 获取 s 与其它参数的交集，结果存入 s
func (s Set) Intersect(sets ...Set) {
	for _, set := range sets {
		for e := range s {
			if _, ok := set[e]; !ok {
				delete(s, e)
			}
		}
	}
}

// 交集（函数）
// 获取所有参数的交集，并返回
func Intersect(sets ...Set) Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取交集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range r {
			if _, ok := set[e]; !ok {
				delete(r, e)
			}
		}
	}
	return r
}

// 补集
// 获取 s 相对于 full 的补集，结果存入 s
func (s Set) Complement(full Set) {
	r := s.Duplicate()
	s.Clear()
	for e := range full {
		if _, ok := r[e]; !ok {
			s[e] = true
		}
	}
}

// 补集（函数）
// 获取 sub 相对于 full 的补集，并返回
func Complement(sub, full Set) Set {
	r := full.Duplicate()
	for e := range sub {
		delete(r, e)
	}
	return r
}