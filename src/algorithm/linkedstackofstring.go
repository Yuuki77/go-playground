package algorithm

type node struct {
	Data string
	Next *node
}

type linkedstackofstring struct {
	head *node
}

func (l *linkedstackofstring) Push(data string) {
	node := &node{Data: data}

	oldHead := l.head
	l.head = node
	l.head.Next = oldHead
}

func (l *linkedstackofstring) Pop() string {
	oldHead := l.head
	l.head = l.head.Next
	return oldHead.Data
}
