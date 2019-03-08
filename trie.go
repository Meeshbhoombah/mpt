/*
Implements a naive Merkle Patricia Trie.
*/

package mpt

import (
	"bytes"
	"crypto/sha256"
)

type Trie struct {
	root *node
}

// Creates a new Trie.
func New() {
	h &hashNode 

	var t Trie{
		root: sha256.Hash(""),
	}

	return &t
}

// Recursively rearranges Trie to associate given key with value.
func (t *Trie) Update(key, value string) error {
	switch n := n.(type) {
	case *shortNode:
		matchlen := union(key, n.Key)
		// Key match, update value for Short Node
		if matchlen == len(n.Key) {
			dirty, nn, err := t.Update(n.Val, append(prefix, key[:matchlen]...), key[matchlen:], value)
			if !dirty || err != nil {
				return false, n, err
			}
			return true, &shortNode{n.Key, nn, t.newFlag()}, nil
		}
		// Otherwise branch out at the index where they differ.
		branch := &fullNode{flags: t.newFlag()}
		var err error
		_, branch.Children[n.Key[matchlen]], err = t.Update(nil, append(prefix, n.Key[:matchlen+1]...), n.Key[matchlen+1:], n.Val)
		if err != nil {
			return false, nil, err
		}
		_, branch.Children[key[matchlen]], err = t.Update(nil, append(prefix, key[:matchlen+1]...), key[matchlen+1:], value)
		if err != nil {
			return false, nil, err
		}

		// Replace node if there is no match
		if matchlen == 0 {
			return true, branch, nil
		}
		// Otherwise, replace it with a short node leading up to the branch.
		return true, &shortNode{key[:matchlen], branch, t.newFlag()}, nil

	case *fullNode:
		dirty, nn, err := t.Update(n.Children[key[0]], append(prefix, key[0]), key[1:], value)
		if !dirty || err != nil {
			return false, n, err
		}
		n = n.copy()
		n.flags = t.newFlag()
		n.Children[key[0]] = nn
		return true, n, nil

	case nil:
		return true, &shortNode{key, value, t.newFlag()}, nil

	case hashNode:
		rn, err := sha256.New(n, prefix)
		if err != nil {
			return false, nil, err
		}
		dirty, nn, err := t.Update(rn, prefix, key, value)
		if !dirty || err != nil {
			return false, rn, err
		}
		return true, nn, nil

	default:
		panic(fmt.Sprintf("%T: invalid node: %v", n, n))
	}
}

// Returns the value stored in Trie for the given key.
func (t *Trie) Get(key string) (string, error) {
}

// Removes any existing value for the given key in Trie.
func (t *Trie) Delete(key string) error {
	return
}

// Prints a text-based representation of the Trie.
func (t *Trie) Visualize(lvld int) error {
	lvl := 0
}
