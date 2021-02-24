package trie_test

import (
	"encoding/json"
	"errors"
	"testing"
	"trie"
)

func TestTrie(t *testing.T) {
	trie := trie.New()
	trie.Write("/", "hello")
	trie.Write("/name", "zhaolu")
	trie.Write("/name/alias", "rongminglu")
	trie.Write("/height", 155.5)
	trie.Write("/height", 155.8)
	trie.Write("age", 22)

	bytes, err := json.Marshal(trie)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	node := trie.Read("/height")
	if node == nil {
		t.Fatal(errors.New("which node should be exists but not"))
	}
	bytes, err = json.Marshal(node)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
