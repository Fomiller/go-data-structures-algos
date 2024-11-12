package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		err := fmt.Errorf("Length of Stack is 0")
		return 0, err
	} else {
		value := s.elements[len(s.elements)-1]

		s.elements = s.elements[:len(s.elements)-1]

		return value, nil
	}
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		err := fmt.Errorf("Length of Stack is 0")
		return 0, err
	} else {
		return s.elements[len(s.elements)-1], nil
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}
