package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.pop()
	for v != nil {
		fmt.Printf("%v->", v)
		v = queue.pop()
	}
}
