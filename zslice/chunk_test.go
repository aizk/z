package zslice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		r   []int
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantS [][]int
	}{
		{
			name:  "success",
			args:  args{
				r:   []int{1, 2, 3, 4, 5},
				n: 2,
			},
			wantS: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:  "success",
			args:  args{
				r:   []int{1, 2, 3, 4},
				n: 2,
			},
			wantS: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:  "success",
			args:  args{
				r:   []int{1, 2, 3, 4, 5},
				n: 100,
			},
			wantS: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name:  "success",
			args:  args{
				r:   []int{1, 2, 3, 4, 5},
				n: 0,
			},
			wantS: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := Chunk[int](tt.args.r, tt.args.n); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Chunk() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
