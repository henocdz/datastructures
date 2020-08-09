package main

import (
	"fmt"

	"github.com/henocdz/datastructures/collections"
)

func linkedList(list *collections.LinkedList) {
	list.InsertAt(1, 0)
	fmt.Println(list.Length)

	list.Push(2)
	list.Push(4)
	list.InsertAt(3, 2)

	fmt.Println(list.Length)
	list.Print()
	fmt.Println("...")
	list.Pop(2)
	list.Print()
	fmt.Println(list.Length)
}

func testListLike(l collections.List, items []int) {
	for _, data := range items {
		l.Push(data)
	}

	l.Print()
}

func main() {
	list := &collections.LinkedList{}
	linkedList(list)
	fmt.Println("<<<<EO LinkedList>>>")

	fmt.Println("Queue expected output: 1,2,3,4")
	queue := &collections.Queue{}
	queueItems := []int{1, 2, 3, 4}
	testListLike(queue, queueItems)
	fmt.Println("\n<<<<EO Queue>>>")

	fmt.Println("Stack expected output: 20,19,18,17,15")
	stack := &collections.Stack{}
	stackItems := []int{15, 16, 17, 18, 19, 20}
	testListLike(stack, stackItems)
}
