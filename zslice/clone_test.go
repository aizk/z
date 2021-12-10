package zslice

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantDst interface{}
	}{
		{
			name:    "[]int copy",
			args:    args{
				src: []int{1, 2, 3},
			},
			wantDst: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := Clone(tt.args.src);

			if !reflect.DeepEqual(dst, tt.wantDst) {
				t.Errorf("Clone() = %v, want %v", dst, tt.wantDst)
			}

			// assert dst is clone
			if reflect.ValueOf(dst).Pointer() == reflect.ValueOf(tt.args.src).Pointer() {
				t.Errorf("Clone() = %v, want %v", dst, tt.wantDst)
			}
		})
	}
}