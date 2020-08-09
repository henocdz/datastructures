package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	//list := &collections.LinkedList{}
	//linkedList(list)
	//fmt.Println("<<<<EO LinkedList>>>")
	//
	//fmt.Println("Queue expected output: 1,2,3,4")
	//queue := &collections.Queue{}
	//queueItems := []int{1, 2, 3, 4}
	//testListLike(queue, queueItems)
	//fmt.Println("\n<<<<EO Queue>>>")
	//
	//fmt.Println("Stack expected output: 20,19,18,17,15")
	//stack := &collections.Stack{}
	//stackItems := []int{15, 16, 17, 18, 19, 20}
	//testListLike(stack, stackItems)

	// TODO: Clean up and add menu, like in school.
	// TODO: Add test cases.
	trie := &collections.Trie{}
	trie.Init()

	dictionary := [...]string{
		"changing", "the", "future", "of",
		"wórk", "futbol", "and", "bot",
		"bots", "boooots", "botcitos", "botsitos",
		"bolsa", "botana", "botica", "botella",
		"bomba",
	}

	for _, p := range dictionary {
		trie.Add(p)
	}

	var term string

	term = "fut"
	fmt.Printf("Contains Prefix {%s}: %t \n", term, trie.ContainsPrefix(term))

	term = "changing"
	fmt.Printf("Contains {%s} %t \n", term, trie.Contains(term))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which word are you looking for?")
	term, _ = reader.ReadString('\n')

	term = strings.Trim(term, "\n")
	matchedWords := trie.Search(term)
	fmt.Printf("Matched words for %s: %v\n", term, matchedWords)
}
