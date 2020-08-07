package collections

import (
	"fmt"
)

/*Node of LinkedList*/
type Node struct {
	Data int
	Next   *Node
}

/*LinkedList type*/
type LinkedList struct {
	Head   *Node
	Length int
}

// preppend add the received Node to the begining of the list
func (l *LinkedList) Preppend(data int) {
	newNode := &Node {
		Data: data,
	}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
	return
}

// append add the received Node to the end of the list
func (l *LinkedList) Push(value int) {
	newNode := &Node{
		Data: value,
	}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		break
	}

	currentNode.Next = newNode
	l.Length++
}

// print prints out LinkedList items
func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

// insertAt adds the given Node into the given position
func (l *LinkedList) InsertAt(data int, position uint) {
	newNode := &Node{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}


	var i uint = 0
	currentNode := l.Head
	prevNode := l.Head
	for currentNode != nil {
		if i == position {
			prevNode.Next = newNode
			newNode.Next = currentNode
			break
		}

		i++
		if currentNode.Next == nil {
			// Index out of range
			panic("Index out of range.")
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}
	l.Length++
}

// remove the Node at given position
func (l *LinkedList) Pop(position uint) int {
	if l.Head == nil {
		panic("Index out of range.")
	}

	var i uint = 0
	currentNode := l.Head
	prevNode := l.Head
	for currentNode != nil {
		if i == position {
			prevNode.Next = currentNode.Next
			break
		}

		i++
		if currentNode.Next == nil {
			// Index out of range
			panic("Index out of range.")
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}
	l.Length--
	return currentNode.Data
}
