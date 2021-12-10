package zslice

import "reflect"

// Clone will clone slice
func Clone(src interface{}) (dst interface{}) {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Slice {
		return
	}

	// copy v slice with value
	vn := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	reflect.Copy(vn, v)
	return vn.Interface()
}

// Up to the official Go SDK 1.11, the simplest way to clone a slice is:
func CloneInt2(s []int) []int {
	return append(s[:0:0], s...)
}

func CloneInt(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return clone
}

