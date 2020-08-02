package main

import (
	"fmt"
)

func main() {

	initialHead := &Node{
		number: 1,
	}

	list := &LinkedList{
		head: initialHead,
	}

	due := &Node{
		number: 2,
	}

	tre := &Node{
		number: 3,
	}

	quattro := &Node{
		number: 4,
	}

	list.append(due)
	list.append(quattro)
	list.insertAt(tre, 2)
	list.print()
	fmt.Println("...")
	list.remove(2)
	list.print()

}
