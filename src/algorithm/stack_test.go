package algorithm

import "testing"

func Test_stack_Push(t *testing.T) {
	type fields struct {
		collection []string
		n          int
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &stack{
				collection: tt.fields.collection,
				n:          tt.fields.n,
			}
			l.Push(tt.args.data)
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		collection []string
		n          int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				collection: []string{"1", "2", "3", "4", "5"},
				n:          5,
			},
			want: "5",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &stack{
				collection: tt.fields.collection,
				n:          tt.fields.n,
			}
			if got := l.Pop(); got != tt.want {
				t.Errorf("stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_resizeArray(t *testing.T) {
	type fields struct {
		collection []string
		n          int
	}
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &stack{
				collection: tt.fields.collection,
				n:          tt.fields.n,
			}
			l.resizeArray(tt.args.size)
		})
	}
}
