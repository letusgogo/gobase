package util

import "testing"

func TestCompareVersions(t *testing.T) {
	type args struct {
		v1 string
		v2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				v1: "2.100.1",
				v2: "2.2.1",
			},
			want: 98,
		},
		{
			name: "t2",
			args: args{
				v1: "2.2.1",
				v2: "2.100.1",
			},
			want: -98,
		},
		{
			name: "t3",
			args: args{
				v1: "2.2.1",
				v2: "2.2.1",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersions(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("CompareVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}
