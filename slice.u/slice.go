package slice_u

import (
	"math/rand"
)

// Append
func Append(s, t []int) {
	s = append(s, t...)
}

// Copy
func Copy(s []int) []int {
	t := make([]int, len(s))
	copy(t, s)
	// or
	t1 := append([]int{}, s...)
	return t1
}

// Cut
func Cut(s []int, i, j int) []int {
	s = append(s[:i], s[j:]...)
	return s
}

// Expand 在中间扩展
func Expand(s []int, index, num int) {
	s = append(s[:index], append(make([]int, num), s[index:]...)...)
}

// Extend 在后面增加
func Extend(s []int, num int)  {
	s = append(s, make([]int, num)...)
}

// Insert 插入一个数据
func Insert(s []int, index, data int)  {
	s = append(s[:index], append([]int{data}, s[index:]...)...)
}

// InsertVector in index
func InsertVector(s, t []int, index int)  {
	s = append(s[:index], append(t, s[index:]...)...)
}

// 弹出首个元素
func POPShift(s []int) int {
	x, s := s[0], s[1:]
	return x
}

// POPBack
func POPBack(s []int) int {
	x, s := s[len(s) - 1], s[:len(s) - 1]
	return x
}

// PUSH
func PUSH(s []int, x int) {
	s = append(s, x)
}

// PUSH FRONT
func PUSHFront(s []int, x int) {
	s = append([]int{x}, s...)
}

// Filtering without allocating
// 不分配新内存过滤
func Filter(s []int) {
	t := s[:0]
	for _, value := range s {
		// 条件过滤参数
		if true {
			t = append(t, value)
		}
	}
}

// Reversing
// To replace the contents of a slice with the same elements but in reverse order:
func Reversing(s []int) {
	for i := len(s)/2 - 1 ; i >= 0; i-- {
		t := len(s) - 1 - i
		s[i], s[t] = s[t], s[i]
	}

	for left, right := 0, len(s) - 1; left < right; left, right = left + 1, right - 1 {
		s[left], s[right] = s[right], s[left]
	}
}

// Shuffling
// 洗牌 打乱顺序
func Shuffling(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}