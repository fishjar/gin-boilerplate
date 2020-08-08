/*
	工具包
*/

package utils

// Index 返回目标字符串 `t` 在 `vs` 中第一次出现位置的索引，
// 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 如果目标字符串 `t` 存在于切片 `vs` 中，则返回 `true`。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 如果切片 `vs` 中的任意一个字符串满足条件 `f`，则返回 `true`。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 如果切片 `vs` 中的所有字符串都满足条件 `f`，则返回 `true`。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个新的切片，新切片由原切片 `vs` 中满足条件 `f` 的字符串构成。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 返回一个新的切片，新切片的字符串由原切片 `vs` 中的字符串经过函数 `f` 映射后得到。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// IFUniqueItem 去重接口
type IFUniqueItem interface {
	GetUniqueIndex() string
}

// RemoveDuplicateElemt 去重函数
func RemoveDuplicateElemt(items []IFUniqueItem) []IFUniqueItem {
	result := make([]IFUniqueItem, 0, len(items))
	temp := map[string]struct{}{}
	for _, item := range items {
		index := item.GetUniqueIndex()
		if _, ok := temp[index]; !ok {
			temp[index] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Union 求并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]struct{}) // slice1 转为 map
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; !ok {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// Intersect 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]struct{}) // slice1 转为 map
	n := make([]string, 0)         // 结果 slice
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			n = append(n, v)
		}
	}
	return n
}

// func Intersect(slice1, slice2 []string) []string {
// 	m := make(map[string]int)
// 	nn := make([]string, 0)
// 	for _, v := range slice1 {
// 		m[v]++
// 	}

// 	for _, v := range slice2 {
// 		times, _ := m[v]
// 		if times == 1 {
// 			nn = append(nn, v)
// 		}
// 	}
// 	return nn
// }

// Difference 求差集，slice1-slice2
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]struct{}) // slice2 转为 map
	n := make([]string, 0)         // 结果 slice
	for _, v := range slice2 {
		m[v] = struct{}{}
	}
	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			n = append(n, v)
		}
	}
	return n
}
