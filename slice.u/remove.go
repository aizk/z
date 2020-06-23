package slice_u

import (
	"unsafe"
	"reflect"
)

func Remove(s interface{}, index ...int) interface{} {
	switch v := s.(type) {
	case []int:
		return RemoveInts(v, index...)
	case []string:
		RemoveStrings(v, index...)
	}
	panic("error type")
}

func RemoveStrings(s []string, index ...int) []string {
	for _, i := range index {
		s = append(s[:i], s[i + 1:]...)
	}
	return s
}

func RemoveInts(s []int, index ...int) []int {
	for _, i := range index {
		s = RemoveInt(s, i)
	}
	return s
}

func RemoveInt(s []int, index int) []int {
	return append(s[:index], s[index + 1:]...)
}

// s 值拷贝
// 传递指针时传递的也是原有指针的拷贝
// 思路：
// 传递指向原有值的指针
// 取得对应的 slice 进行 append 操作
// 产生新的 slice
// 将新的和旧的都转化成 SliceHeader
// 将新的值赋给旧的内存
func RemoveInt2(s *[]int, index int) {
	// 产生新的值
	v := *s
	// 共享底层数组
	newSlice := append(v[:index], v[index + 1:]...)
	// 修改不了 s 指针指向的位置？
	// 只修改了指针拷贝的值！
	//s = &v
	ns := (*reflect.SliceHeader)(unsafe.Pointer(&newSlice))
	os := (*reflect.SliceHeader)(unsafe.Pointer(s))
	os.Data = ns.Data
	os.Len = ns.Len
	os.Cap = ns.Cap
}

func RemoveInt3(s []int, index int) {
	copy(s[index:], s[index + 1:])
	s[len(s) - 1] = 0
	s = s[:len(s) - 1]
}

// Delete without preserving order 删除不保留顺序
func RemoveWithoutOrder(s []int, index int) []int {
	// 将最后一个元素赋值给当前元素
	s[index] = s[len(s) - 1]
	// 删除最后一个元素
	s = s[:len(s) - 1]
	return s
}
