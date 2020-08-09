package collections

import "fmt"

type Queue struct {
	Head *BasicNode
	Tail *BasicNode
	Length int
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue) Push(data int) {
	newNode := &BasicNode{
		Data: data,
	}

	if q.Head == nil {
		q.Head = newNode
	}

	if q.Tail != nil {
		q.Tail.Next = newNode
	}
	q.Tail = newNode
	q.Length++
}

func (q *Queue) Pop() int {
	if q.Head == nil {
		panic("Empty queue.")
	}

	data := q.Head.Data

	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}

	q.Length--
	return data
}

func (q *Queue) Print() {
	for !q.IsEmpty() {
		fmt.Printf("%d", q.Pop())
		if !q.IsEmpty() {
			fmt.Print(" >> ")
		}
	}
}