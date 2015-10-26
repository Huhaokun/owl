package utils

import (
	"container/list"
)

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{l}
}

func (s *Stack) Pop() interface{} {
	res := s.l.Back()
	if e != nil {
		s.l.Remove(e)
		return e
	}
	return nil
}

func (s *Stack) Push(val interface{}) {
	s.l.PushBack(val)
}

func (s *Stack) Empty() bool {
	if s.l.Len() == 0 {
		return true
	} else {
		return false
	}
}
