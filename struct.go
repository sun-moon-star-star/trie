package trie

import "strings"

type NodeData struct {
	Key   string
	Value interface{}
}

type Comparator func(*string, *string) bool // equal

type Node struct {
	Data NodeData

	Father   *Node `json:"-"`
	Next     *Node
	Children *Node
}

type Trie struct {
	Depth    uint32
	ValueCnt uint32 // count(Node.Data.Value != nil)
	Cnt      uint32

	Comparator Comparator    `json:"-"`
	SplitKey   SplitStrategy `json:"-"`
	Root       Node
}

type SplitStrategy func(string) []string

type Options struct {
	RootKey   string
	RootValue interface{}

	Comparator Comparator    `json:"-"`
	SplitKey   SplitStrategy `json:"-"`
}

var DefaultOptions = Options{
	RootKey:   "",
	RootValue: nil,
	Comparator: func(x, y *string) bool {
		if x == nil || y == nil {
			return x == y
		}
		return x == y
	},
	SplitKey: func(key string) []string {
		return strings.Split(key, "/")
	},
}
