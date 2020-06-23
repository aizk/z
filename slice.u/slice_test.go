package slice_u

import "testing"

func TestRemove(t *testing.T) {
	x := []int{1, 2, 3}
	x = Remove(x, 1).([]int)

	//x = RemoveInt(x, 1)
	t.Log(x)

	//Remove(x, 1)
	//t.Log(x)
	//Remove(x, 1)
	//t.Log(x)
	t.Log(append(x[:1], x[2:]...))
}
