package collections

import (
	"fmt"
)

/*LinkedList type*/
type LinkedList struct {
	Head   *BasicNode
	Length int
}

// prepend add the received BasicNode to the beginning of the list
func (l *LinkedList) Prepend(data int) {
	newNode := &BasicNode{
		Data: data,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Length++
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
	return
}

// IsEmpty
func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

// append add the received BasicNode to the end of the list
func (l *LinkedList) Push(data int) {
	newNode := &BasicNode{
		Data: data,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Length++
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

// insertAt adds the given BasicNode into the given position
func (l *LinkedList) InsertAt(data int, position uint) {
	newNode := &BasicNode{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Length++
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

// remove the BasicNode at given position
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
