/*
Implements a naive Merkle Patricia Trie.
*/

package mpt

type Trie struct {
	root *node
}

// Creates a new Trie.
func (t *Trie) New() {
	return
}

// Rearranges Trie to associate given key with value.
func (t *Trie) Update(key, value string) error {
	return
}

// Returns the value stored in Trie for the given key.
func (t *Trie) Get(key string) (string, error) {
	return
}

// Removes any existing value for the given key in Trie.
func (t *Trie) Delete(key string) error {
	return
}

// Print string representation of Trie.
func (t *Trie) Print(lvld int) (string, error) {
	lvl := 0
}
