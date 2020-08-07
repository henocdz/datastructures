package main

import (
	"fmt"

	"github.com/henocdz/datastructures/collections"
)

func main() {
	list := &collections.LinkedList{}
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
