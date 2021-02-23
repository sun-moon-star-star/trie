package trie

type NodeData struct {
	Key   string
	Value interface{}
}

type Comparator func(*string, *string) bool // equal

type Node struct {
	Data NodeData

	Father   *Node
	Next     *Node
	Children *Node
}

type Trie struct {
	Depth      uint32
	Comparator Comparator
	Root       Node
}

type Options struct {
	RootKey   string
	RootValue interface{}

	Comparator Comparator
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
}
