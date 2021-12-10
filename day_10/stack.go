package main

import(
	"container/list"
)

type Stack struct {
	values *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) Add(element rune) {
	q.values.PushFront(element)
}

func (q *Stack) Peek() rune {
	return q.values.Front().Value.(rune)
}

func (q *Stack) Pop() rune {
	value := q.values.Front().Value
	q.values.Remove(q.values.Front())
	return value.(rune)
}

func (q *Stack) Size() int {
	return q.values.Len()
}