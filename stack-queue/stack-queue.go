package stackqueue

type Stack struct {
	s []uint
}

type Queue struct {
	q []uint
}

func NewStack() *Stack {
	return &Stack{make([]uint, 0)}
}

func NewQueue() *Queue {
	return &Queue{make([]uint, 0)}
}

func (s *Stack) Len() int {
	return len(s.s)
}

func (q *Queue) Len() int {
	return len(q.q)
}

func (s *Stack) Peek() (uint, bool) {
	if len(s.s) == 0 {
		return 0, false
	}
	return s.s[len(s.s) - 1], true
}

func (q *Queue) Peek() (uint, bool) {
	if len(q.q) == 0 {
		return 0, false
	}
	return q.q[0], true
}

func (s *Stack) Pop() (uint, bool) {
	if len(s.s) == 0 {
		return 0, false
	}
	v := s.s[len(s.s) - 1]
	s.s = s.s[:len(s.s)-1]
	return v, true
}

func (q *Queue) Pop() (uint, bool) {
	if len(q.q) == 0 {
		return 0, false
	}
	v := q.q[0]
	q.q = q.q[1:]
	return v, true
}


func (s *Stack) Push(v uint) {
	s.s = append(s.s, v)
}

func (q *Queue) Push(v uint) {
	q.q = append(q.q, v)
}

