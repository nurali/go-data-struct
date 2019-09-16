package singly

type Singly interface {
	Insert(no int)
	Delete(no int)
	ToArray() []int
}

type node struct {
	data int
	next *node
}

type singly struct {
	start *node
}

func NewSingly(numbers ...int) Singly {
	singly := &singly{}
	for _, no := range numbers {
		singly.Insert(no)
	}
	return singly
}

func (s *singly) Insert(no int) {
	if s.start == nil {
		s.start = &node{data: no}
		return
	}

	newNode := &node{data: no}
	last := s.last()
	last.next = newNode
}

func (s *singly) Delete(no int) {
	if s.start == nil {
		return
	}

	prev := s.start
	for n := s.start; n != nil; n = n.next {
		if n.data == no {
			if prev.next == n.next {
				s.start = n.next
				return
			}
			prev.next = n.next
			return
		}
		prev = n
	}
}

func (s *singly) ToArray() []int {
	if s.start == nil {
		return []int{}
	}

	var arr []int
	for n := s.start; n != nil; n = n.next {
		arr = append(arr, n.data)
	}
	return arr
}

func (s *singly) last() *node {
	for n := s.start; ; n = n.next {
		if n.next == nil {
			return n
		}
	}
}
