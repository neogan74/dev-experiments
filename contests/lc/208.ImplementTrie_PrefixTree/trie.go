package trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (tr *Trie) Insert(word string) {
	node := tr
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (tr *Trie) Search(word string) bool {
	node := tr.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (tr *Trie) StartsWith(prefix string) bool {
	node := tr.SearchPrefix(prefix)
	return node != nil
}

func (tr *Trie) SearchPrefix(s string) *Trie {
	node := tr
	for _, c := range s {
		idx := c - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
