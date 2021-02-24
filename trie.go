package trie

func NewWithOptions(options Options) *Trie {
	trie := &Trie{
		Depth:      0,
		ValueCnt:   0,
		Cnt:        1,
		Comparator: options.Comparator,
		SplitKey:   options.SplitKey,
		IgnoreKey:  options.IgnoreKey,
	}

	if trie.Comparator == nil {
		trie.Comparator = DefaultOptions.Comparator
	}

	if trie.SplitKey == nil {
		trie.SplitKey = DefaultOptions.SplitKey
	}

	if trie.IgnoreKey == nil {
		trie.IgnoreKey = DefaultOptions.IgnoreKey
	}

	return trie
}

func New() *Trie {
	return NewWithOptions(DefaultOptions)
}

func (trie *Trie) matchChildrenNode(node *Node, key *string) *Node {
	if node == nil || node.Children == nil {
		return nil
	}

	for ch := node.Children; ch != nil; ch = ch.Next {
		if trie.Comparator(&ch.Data.Key, key) {
			return ch
		}
	}

	return nil
}

func (trie *Trie) Write(key string, value interface{}) uint32 {
	keys := trie.SplitKey(key)
	len := len(keys)
	if len == 0 {
		return 0
	}

	cur, cnt, depth := &trie.Header, uint32(0), uint32(0)

	for idx := 0; idx < len; idx++ {
		if trie.IgnoreKey(keys[idx]) {
			continue
		}
		depth++
		ch := trie.matchChildrenNode(cur, &keys[idx])
		if ch == nil {
			ch = &Node{}
			ch.Data.Key = keys[idx]
			ch.Father = cur
			ch.Next = cur.Children
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
	keys := trie.SplitKey(key)
	len := len(keys)

	if len == 0 {
		return nil
	}

	cur := &trie.Header

	for idx := 0; idx < len && cur != nil; idx++ {
		if !trie.IgnoreKey(keys[idx]) {
			cur = trie.matchChildrenNode(cur, &keys[idx])
		}
	}

	return cur
}
