package main

import (
	"fmt"
)

/*Node of LinkedList*/
type Node struct {
	number int
	next   *Node
}

/*LinkedList type*/
type LinkedList struct {
	head *Node
}

// preppend add the received Node to the begining of the list
func (l *LinkedList) preppend(n *Node) {
	if l.head == nil {
		l.head = n
		return
	}

	n.next = l.head
	l.head = n
	return
}

// append add the received Node to the end of the list
func (l *LinkedList) append(n *Node) {
	if l.head == nil {
		l.head = n
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
		break
	}

	currentNode.next = n
}

// print prints out LinkedList items
func (l *LinkedList) print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.number)
		currentNode = currentNode.next
	}
}

// insertAt adds the given Node into the given position
func (l *LinkedList) insertAt(n *Node, position uint) {
	if l.head == nil {
		panic("Index out of range.")
	}

	var i uint = 0
	currentNode := l.head
	prevNode := l.head
	for currentNode != nil {
		if i == position {
			prevNode.next = n
			n.next = currentNode
			break
		}

		i++
		if currentNode.next == nil {
			// Index out of range
			panic("Index out of range.")
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
}

// remove the Node at given position
func (l *LinkedList) remove(position uint) {
	if l.head == nil {
		panic("Index out of range.")
	}

	var i uint = 0
	currentNode := l.head
	prevNode := l.head
	for currentNode != nil {
		if i == position {
			prevNode.next = currentNode.next
			break
		}

		i++
		if currentNode.next == nil {
			// Index out of range
			panic("Index out of range.")
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
}
