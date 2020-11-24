package util

import "testing"

func Test_isInteger(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ok", args: args{a: 11}, want: true},
		{name: "failed", args: args{a: 11.2}, want: false},
		{name: "ok1", args: args{a: 0.0}, want: true},
		{name: "failed1", args: args{a: 0.1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInteger(tt.args.a); got != tt.want {
				t.Errorf("isInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBoolean(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "str-failed", args: args{val: "n"}, want: false, wantErr: true},
		{name: "str-ok", args: args{val: "true"}, want: true, wantErr: false},
		{name: "str-ok", args: args{val: "false"}, want: false, wantErr: false},

		{name: "str-failed", args: args{val: "3"}, want: false, wantErr: true},
		{name: "str-ok", args: args{val: "1"}, want: true, wantErr: false},
		{name: "str-ok", args: args{val: "0"}, want: false, wantErr: false},

		{name: "int-failed", args: args{val: 3}, want: false, wantErr: true},
		{name: "int-ok", args: args{val: 1}, want: true, wantErr: false},
		{name: "int-ok", args: args{val: 0}, want: false, wantErr: false},

		{name: "float-failed", args: args{val: 0.3}, want: false, wantErr: true},
		{name: "float-ok", args: args{val: 1.0}, want: true, wantErr: false},
		{name: "float-ok", args: args{val: 0.0}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBoolean(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBoolean() got = %v, want %v", got, tt.want)
			}
		})
	}
}
