package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()
	books := [5]string{"꼬북좌 맥심 화보집", "어린왕자", "용의자 X의 헌신", "백야행", "바보 빅터"}
	for i := 0; i < 5; i++ {
		stack.Push(books[i])
	}
	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v=>", val)
		val = stack.Pop()
	}
}
