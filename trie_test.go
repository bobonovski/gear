package gear

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	words := []string{"hello", "yes", "hi", "yikes", "oops", "ouch", "zoo", "bingo"}
	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}
	assert.Equal(t, true, trie.Exists("yes"))

	trie.Delete("zoo")
	assert.Equal(t, false, trie.Exists("zoo"))

	ws := trie.FindWithPrefix("o")
	sort.Strings(ws)
	assert.Equal(t, []string{"oops", "ouch"}, ws)

	ws = trie.FindWithPrefix("b")
	assert.Equal(t, []string{"bingo"}, ws)
}
