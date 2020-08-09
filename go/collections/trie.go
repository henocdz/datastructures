package collections

import (
	"unicode"
)

type Trie struct {
	children map[rune]*TrieNode
}

func (t *Trie) Init() {
	t.children = make(map[rune]*TrieNode)
}

func (t *Trie) Add(word string) {
	wr := []rune(word)
	wordLength := len(wr) - 1
	currentChildren := t.children

	for cpos, c := range wr {
		r := unicode.ToLower(c)
		nextNode, ok := currentChildren[r]
		var wordLastNode *TrieNode
		if ok {
			currentChildren = nextNode.Children
			wordLastNode = nextNode
		} else {
			newNode := &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
			wordLastNode = newNode
			currentChildren[r] = newNode
			currentChildren = newNode.Children
		}

		if cpos == wordLength {
			wordLastNode.endOfWord = true
		}
	}
}

func (t *Trie) findPrefixNode(term string) *TrieNode {
	wr := []rune(term)
	currentChildren := t.children

	var wordLastNode *TrieNode
	for _, c := range wr {
		nextNode, ok := currentChildren[c]
		if !ok {
			return nil
		}
		currentChildren = nextNode.Children
		wordLastNode = nextNode
	}
	return wordLastNode
}

func (t *Trie) ContainsPrefix(term string) bool {
	wr := []rune(term)
	matches := false

	currentChildren := t.children
	for _, c := range wr {
		nextNode, ok := currentChildren[c]
		if ok {
			currentChildren = nextNode.Children
			matches = true
		} else {
			matches = false
			break
		}
	}

	return matches
}

func (t *Trie) Contains(term string) bool {
	prefixNode := t.findPrefixNode(term)
	if prefixNode != nil {
		return prefixNode.endOfWord
	}
	return false
}

func (t *Trie) findNodeWords(node *TrieNode, accRunes []rune, matchedWords [][]rune) [][]rune {
	if node.endOfWord {
		matchedWords = append(matchedWords, accRunes)
	}

	for wordRune, child := range node.Children {
		prevRunes := make([]rune, len(accRunes))
		copy(prevRunes, accRunes) // Pretty sure this is not performant at all :thinking:
		prevRunes = append(prevRunes, wordRune)
		matchedWords = t.findNodeWords(child, prevRunes, matchedWords)
	}

	return matchedWords
}

func (t *Trie) Search(term string) []string {
	prefixNode := t.findPrefixNode(term)
	var matchedWordsString []string

	if prefixNode == nil {
		return matchedWordsString
	}

	prevRunes := []rune(term)
	matchedWords := make([][]rune, 0)

	matchedWords = t.findNodeWords(prefixNode, prevRunes, matchedWords)

	for _, wordRunes := range matchedWords {
		matchedWordsString = append(matchedWordsString, string(wordRunes))
	}

	return matchedWordsString
}
