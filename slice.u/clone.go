package slice_u

// clone slice
func Clone(s interface{}) (r interface{}) {
	switch v := s.(type) {
	case []int:
		//r = CloneInt(v)
		return CloneInt(v)
	}
	return
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

