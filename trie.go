package trie

func NewWithOptions(options Options) *Trie {
	return &Trie{
		Depth:      1,
		Comparator: options.Comparator,
		Root: Node{
			Data: NodeData{
				Key:   options.RootKey,
				Value: options.RootValue,
			},

			Father:   nil,
			Next:     nil,
			Children: nil,
		},
	}
}

func New() *Trie {
	return NewWithOptions(DefaultOptions)
}

func (trie *Trie) matchChildrenNode(fa *Node, key *string) *Node {
	if fa == nil || fa.Children == nil {
		return nil
	}

	ch := fa.Children
	for ch != nil {
		if trie.Comparator(&fa.Data.Key, key) {
			return ch
		}
		ch = ch.Next
	}

	return nil
}

func (trie *Trie) Set(keys []string, value interface{}) int32 {
	cur := &trie.Root
	cnt := int32(0)

	for _, key := range keys {
		ch := trie.matchChildrenNode(cur, &key)
		if ch == nil {
			ch = &Node{
				Data: NodeData{
					Key:   key,
					Value: nil,
				},

				Father:   cur,
				Next:     cur.Children,
				Children: nil,
			}
			cnt++
		}
		cur = ch
	}
	cur.Data.Value = value

	return cnt
}

func (trie *Trie) Get(keys []string) *Node {
	cur := &trie.Root

	for _, key := range keys {
		ch := trie.matchChildrenNode(cur, &key)
		if ch == nil {
			return nil
		}
		cur = ch
	}

	return cur
}
