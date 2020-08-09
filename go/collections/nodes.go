package collections

/*BasicNode for: LinkedList, Queue, Stack */
type BasicNode struct {
	Data int
	Next *BasicNode
}

type TrieNode struct {
	Children  map[rune]*TrieNode
	endOfWord bool
}
