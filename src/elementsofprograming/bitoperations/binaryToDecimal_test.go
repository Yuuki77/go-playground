package elementsofprograming

import "testing"

func Test_binaryToDecimal(t *testing.T) {
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
				x: 111111,
			},
			want: 63,
		},
		{
			name: "test2",
			args: args{
				x: 101,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryToDecimal(tt.args.x); got != tt.want {
				t.Errorf("binaryToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
