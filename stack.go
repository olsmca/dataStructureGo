package datastructurego

import "fmt"

type Stack []string

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", true
	} else {
		index := s.Size() - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, false
	}
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) ToString() string {
	return fmt.Sprintf("Stack: %v", *s)
}
