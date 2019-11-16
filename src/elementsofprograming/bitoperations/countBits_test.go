package elementsofprograming

import (
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				x: 100,
			},
			want: 3,
		},
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				x: 255,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.x); got != tt.want {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitAnd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				x: 45,
				y: 25,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitAnd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("bitAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
