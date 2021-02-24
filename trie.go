package trie

func NewWithOptions(options Options) *Trie {
	trie := &Trie{
		Depth:      1,
		ValueCnt:   0,
		Cnt:        1,
		Comparator: options.Comparator,
		SplitKey:   options.SplitKey,
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

	if trie.Comparator == nil {
		trie.Comparator = DefaultOptions.Comparator
	}

	if trie.SplitKey == nil {
		trie.SplitKey = DefaultOptions.SplitKey
	}

	if trie.Root.Data.Value != nil {
		trie.ValueCnt = 1
	}

	return trie
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

func (trie *Trie) Write(key string, value interface{}) uint32 {
	keys := trie.SplitKey(key)
	len := len(keys)
	var depth uint32

	if len == 0 {
		return 0
	}

	cur := &trie.Root
	cnt := uint32(0)
	idx := 0

	if keys[0] == "" {
		idx = 1
		depth = 1
	}

	for idx < len {
		key := keys[idx]
		idx++
		if key == "" {
			continue
		}
		depth++
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
			cur.Children = ch
			cnt++
		}
		cur = ch
	}

	if cur.Data.Value == nil && value != nil {
		trie.ValueCnt++
	}
	if cur.Data.Value != nil && value == nil {
		trie.ValueCnt--
	}

	cur.Data.Value = value

	if depth > trie.Depth {
		trie.Depth = depth
	}

	trie.Cnt += cnt
	return cnt
}

func (trie *Trie) Read(key string) *Node {
	cur := &trie.Root
	keys := trie.SplitKey(key)

	for _, key := range keys {
		ch := trie.matchChildrenNode(cur, &key)
		if ch == nil {
			return nil
		}
		cur = ch
	}

	return cur
}
