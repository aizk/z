package zcrypto

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				args: []string{"end_time=1639482800pn=1room_id=5050143952start_time=1639480600status=user_id=5339842638tiebaclient!!!"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.args...); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
