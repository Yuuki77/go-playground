package algorithm

type stack struct {
	collection []string
	n          int
}

func (l *stack) Push(data string) {
	if l.n == len(l.collection) {
		l.resizeArray(l.n * 2)
	}
	l.collection[l.n] = data
	l.n++
}
func (l *stack) Pop() string {
	l.n--
	item := l.collection[l.n]
	if l.n > 0 && l.n == 4/len(l.collection) {
		l.resizeArray(len(l.collection) / 2)
	}
	return item
}

func (l *stack) resizeArray(size int) {
	temp := make([]string, size)

	for index := 0; index < l.n; index++ {
		temp[index] = l.collection[index]
	}

	l.collection = temp
}
