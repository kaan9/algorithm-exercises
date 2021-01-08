package linkedlist

type node struct {
	val  uint
	next *node
}

type List struct {
	head *node
}

func (l *List) peek() {
	return l.head.val
}

func (l *List) prepend(v uint) {
	n := &node{v, l.head}
	l.head = n
}

func (l *List) append(v uint) {
	n := &node{v, nil}
	next := &l.head
	for (*next) != nil {
		next = &(*next).next
	}
	*next = n
}

//TODO: complete this implementation
