package Trie

type Node[T any] struct {
	last     bool
	parent   *Node[T]
	children map[int32]*Node[T]
}

type Trie[T any] struct {
	root *Node[T]
}

func (trie *Trie[T]) Init() {
	trie.root = &Node[T]{children: map[int32]*Node[T]{}}
}

func (trie Trie[T]) Add(item string) {
	currentNode := trie.root
	for _, r := range item {
		if _, ok := currentNode.children[r]; ok {
			currentNode = currentNode.children[r]
		} else {
			node := &Node[T]{children: map[int32]*Node[T]{}, parent: currentNode}
			currentNode.children[r] = node
			currentNode = node
		}
	}
	currentNode.last = true
}

func (trie Trie[T]) Search(item string) bool {
	currentNode := trie.root
	for _, r := range item {
		if _, ok := currentNode.children[r]; ok {
			currentNode = currentNode.children[r]
		} else {
			return false
		}
	}
	if currentNode.last == false {
		return false
	}
	return true
}

func (trie Trie[T]) Remove(item string) bool {
	currentNode := trie.root
	for _, r := range item {
		if _, ok := currentNode.children[r]; ok {
			currentNode = currentNode.children[r]
		} else {
			return false
		}
	}
	if currentNode.last == false {
		return false
	}
	currentNode.last = false
	symbolIndex := len(item) - 1
	for len(currentNode.children) == 0 {
		delete(currentNode.children, int32(item[symbolIndex]))
		currentNode = currentNode.parent
		symbolIndex--
	}
	return true
}
