package algorithm

import (
	"testing"
)

func Test_linkedstackofstring_Pop(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		l := &linkedstackofstring{}

		for index := 0; index <= 100; index++ {
			l.Push(string(index))
		}
		for index := 100; index >= 0; index-- {
			ans := l.Pop()
			if ans != string(index) {
				t.Errorf("push ans pop = %v, want %v", ans, string(index))
			}

		}
		// if got := l.Pop(); !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("linkedstackofstring.Pop() = %v, want %v", got, tt.want)
		// }
	})
}
