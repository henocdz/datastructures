package collections

import "fmt"

type Stack struct {
	Head *BasicNode
}

func (s *Stack) Push(data int) {
	newNode := &BasicNode{
		Data: data,
		Next: s.Head,
	}

	s.Head = newNode
}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Empty stack.")
	}

	data := s.Head.Data
	s.Head = s.Head.Next
	return data
}

func (s *Stack) Print() {
	for !s.IsEmpty() {
		fmt.Printf("%d", s.Pop())
		if !s.IsEmpty() {
			fmt.Print(" > ")
		}
	}
}
