package zslice

func EqualSlice(x, y []interface{}) bool {
	if len(x) != len(y) {
		return false
	}

	for _, v1 := range x {
		find := false
		for _, v2 := range y {
			if v1 == v2 {
				find = true
				break
			}
		}
		if !find {
			return false
		}
	}
	return true
}
