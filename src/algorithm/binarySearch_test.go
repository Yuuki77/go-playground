package algorithm

import (
	"math"
	"testing"
)

func Test_binarySearchError(t *testing.T) {
	s := make([]int8, math.MaxInt8)
	for index := 0; index < len(s); index++ {
		s[index] = int8(index)
	}
	type args struct {
		array  []int8
		lo     int8
		hi     int8
		target int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				array:  s,
				lo:     0,
				hi:     int8(len(s) - 1),
				target: math.MaxInt8 - 1,
			},
			want: "integer overflow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := binarySearchError(tt.args.array, tt.args.lo, tt.args.hi, tt.args.target); got.Error() != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	s := make([]int8, math.MaxInt8)
	for index := 0; index < len(s); index++ {
		s[index] = int8(index)
	}
	type args struct {
		array  []int8
		lo     int8
		hi     int8
		target int8
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array:  s,
				lo:     0,
				hi:     int8(len(s) - 1),
				target: math.MaxInt8 - 1,
			},
			want:    int8(len(s) - 1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binarySearch(tt.args.array, tt.args.lo, tt.args.hi, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("binarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
