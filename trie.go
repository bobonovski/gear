package gear

type Trie struct {
	root *node
}

type node struct {
	// whether this node is the end of a word
	end      bool
	children map[string]*node
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{
			end:      false,
			children: make(map[string]*node),
		},
	}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
	chars := t.toChars(word)
	r := t.root
	for i := 0; i < len(chars); i++ {
		if _, ok := r.children[chars[i]]; !ok {
			r.children[chars[i]] = &node{
				end:      false,
				children: make(map[string]*node),
			}
		}
		if i == len(chars)-1 {
			r.children[chars[i]].end = true
		}
		r = r.children[chars[i]]
	}
}

// Delete deletes the word from the trie if the word // exists or do nothing.
func (t *Trie) Delete(word string) {
	n := t.find(word)
	if n.end == true {
		n.end = false
	}
}

// Exists check the existence of the word.
func (t *Trie) Exists(word string) bool {
	n := t.find(word)
	if n.end == true {
		return true
	}
	return false
}

// FindWithPrefix returns all words that have the prefix.
func (t *Trie) FindWithPrefix(prefix string) []string {
	words := make([]string, 0)
	n := t.find(prefix)
	if n == nil {
		return words
	}
	suffix := t.getSuffix(n)
	for _, s := range suffix {
		words = append(words, prefix+s)
	}
	return words
}

// Get the suffix start from the node.
func (t *Trie) getSuffix(n *node) []string {
	var suffix []string
	for k, v := range n.children {
		sx := t.getSuffix(v)
		if len(sx) > 0 {
			for _, s := range sx {
				suffix = append(suffix, k+s)
			}
		} else {
			suffix = append(suffix, k)
		}
	}
	return suffix
}

// Find the node which is the end of the word.
func (t *Trie) find(word string) *node {
	chars := t.toChars(word)
	r := t.root
	for i := 0; i < len(chars); i++ {
		if _, ok := r.children[chars[i]]; !ok {
			break
		}
		r = r.children[chars[i]]
		if i == len(chars)-1 {
			return r
		}
	}
	return nil
}

// Convert word to character slice.
func (t *Trie) toChars(word string) []string {
	var chars []string
	for _, c := range word {
		chars = append(chars, string(c))
	}
	return chars
}
