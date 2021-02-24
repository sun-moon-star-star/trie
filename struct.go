package trie

import "strings"

type NodeData struct {
	Key   string
	Value interface{}
}

type Node struct {
	Data NodeData

	Father   *Node `json:"-"`
	Next     *Node
	Children *Node
}

type Comparator func(*string, *string) bool // NodeData.Key

type SplitStrategy func(string) []string

type IgnoreStrategy func(string) bool

type Trie struct {
	Depth    uint32
	ValueCnt uint32 // count(Node.Data.Value != nil)
	Cnt      uint32

	Comparator Comparator     `json:"-"`
	SplitKey   SplitStrategy  `json:"-"`
	IgnoreKey  IgnoreStrategy `json:"-"`

	Header Node // no use
}

type Options struct {
	Comparator Comparator     `json:"-"`
	SplitKey   SplitStrategy  `json:"-"`
	IgnoreKey  IgnoreStrategy `json:"-"`
}

var DefaultOptions = Options{
	Comparator: func(x, y *string) bool {
		if x == nil || y == nil {
			return x == y
		}
		return *x == *y
	},
	SplitKey: func(key string) []string {
		return strings.Split(key, "/")
	},
	IgnoreKey: func(key string) bool {
		return key == ""
	},
}
