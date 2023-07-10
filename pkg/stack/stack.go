package stack

import "errors"

type Stack[T interface{}] struct {
	list []T
}

func (s *Stack[T]) Len() int {
	return len(s.list)
}

func (s *Stack[T]) Push(item T) {
	s.list = append(s.list, item)
}

func (s *Stack[T]) Pop() (T, error) {
	popIdx := s.Len() - 1

	if s.Len() == 0 {
		err := errors.New("stack is empty")
		var nullValue T
		return nullValue, err
	}

	lastItem := s.list[popIdx]
	s.list = s.list[:popIdx]

	return lastItem, nil
}

func (s *Stack[T]) Strategy(f func(idx int, item T)) {
	for idx, item := range s.list {
		f(idx, item)
	}
}
