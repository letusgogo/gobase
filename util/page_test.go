package util

import (
	"reflect"
	"testing"
)

func TestPageSliceOle(t *testing.T) {
	emptySlice := make([]string, 0)
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	type args struct {
		slice    interface{}
		pageable Pageable
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "testEmptySlice",
			args: args{
				slice:    emptySlice,
				pageable: NewDefaultPage(0, 1),
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "test(1,5)Slice",
			args: args{
				slice:    slice,
				pageable: NewDefaultPage(1, 5),
			},
			want:    []int{0, 1, 2, 3, 4},
			wantErr: false,
		},
		{
			name: "test(2,5)Slice",
			args: args{
				slice:    slice,
				pageable: NewDefaultPage(2, 5),
			},
			want:    []int{5, 6, 7, 8, 9},
			wantErr: false,
		},
		{
			name: "test(1,5)Slice",
			args: args{
				slice:    slice,
				pageable: NewDefaultPage(1, 5),
			},
			want:    []int{0, 1, 2, 3, 4},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PageSliceOle(tt.args.slice, tt.args.pageable)
			if (err != nil) != tt.wantErr {
				t.Errorf("PageSliceOle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageSliceOle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageSlice(t *testing.T) {
	type args[T any] struct {
		slice    []T
		pageable Pageable
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "test(1,5)Slice",
			args: args[int]{
				slice:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				pageable: NewDefaultPage(1, 5),
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "test(2,5)Slice",
			args: args[int]{
				slice:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				pageable: NewDefaultPage(2, 5),
			},
			want: []int{5, 6, 7, 8, 9},
		},
		{
			name: "test(3,5)Slice",
			args: args[int]{
				slice:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				pageable: NewDefaultPage(3, 5),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PageSlice(tt.args.slice, tt.args.pageable)
			if (err != nil) != tt.wantErr {
				t.Errorf("PageSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
