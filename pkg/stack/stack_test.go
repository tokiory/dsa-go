package stack_test

import (
	"algo/pkg/stack"
	"testing"
)

func TestStack_Len(t *testing.T) {
	s := stack.Stack[int]{}
	const EmptyStackLength int = 0

	if s.Len() != EmptyStackLength {
		t.Fail()
	}
}

func TestStack_Push(t *testing.T) {
	s := stack.Stack[int]{}
	const ExpectedLength = 5
	list := make([]int, 0, 5)

	for i := 0; i < ExpectedLength; i++ {
		s.Push(i * 10)
	}

	s.Strategy(func(_, item int) {
		list = append(list, item)
	})

	if len(list) != ExpectedLength {
		t.Fail()
	}
}

func TestStack_Pop(t *testing.T) {
	s := stack.Stack[int]{}
	const PushAmount = 5
	const PopAmount = 3
	const ExpectedLength = PushAmount - PopAmount

	list := make([]int, 0, 5)

	for i := 0; i < PushAmount; i++ {
		s.Push(i * 10)
	}

	for i := 0; i < PopAmount; i++ {
		_, err := s.Pop()

		if err != nil {
			t.Fatal(err)
		}
	}

	s.Strategy(func(_, item int) {
		list = append(list, item)
	})

	if len(list) != ExpectedLength {
		t.Fail()
	}
}
