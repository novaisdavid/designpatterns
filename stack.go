package main

type Stack struct {
	empety bool
	size   int
}

func NewStack() Stack {
	return Stack{empety: true, size: 0}
}

func (s Stack) EmpetyStack() bool {
	return s.empety
}

func (s *Stack) Push(a interface{}) {
	s.size = s.size + 1
	s.empety = false
}

func (s *Stack) Pop() {
	s.size = s.size - 1

	if s.size == 0 {
		s.empety = true
	}
}
