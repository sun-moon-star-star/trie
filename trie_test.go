package trie_test

import (
	"encoding/json"
	"testing"
	"trie"
)

func TestTrie(t *testing.T) {
	trie := trie.New()
	trie.Write("/", "hello")
	trie.Write("/name", "zhaolu")
	trie.Write("/name/alias", "rongminglu")
	trie.Write("/age", 21)
	trie.Write("/height", 155.5)

	bytes, err := json.Marshal(trie)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
