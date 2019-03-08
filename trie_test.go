package mpt

import (
	"crypto/sha256"
	"testing"
	"fmt"
)

func TestCreateNewTrie(t *testing.T) {
	trie := New()

	want := sha256.New("")

	if t.root != want {
		t.Errorf("Empty trie has root %s, want %s", t.root, want)
	}
}

func TestUpdateTrie(t *testing.T) {
	var updateTrieWithCases = []struct {
		k string
		v string

		w bool
	}{
		{"dog", "cat", true},
		{"sad", "happy", true},
		{"alan", "cs", true},
	}

	for _, c := range updateTrieWithCases {
		err := Trie.Update(c.k, c.v) {
			fmt.Errorf("Error updating trie: %s", err)
		}	
	}
}
