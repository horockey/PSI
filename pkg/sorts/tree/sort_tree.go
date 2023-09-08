package tree

type node struct {
	key   int
	left  *node
	right *node
}

type sortTree struct {
	root *node
}

func (t *sortTree) insert(el int) {
	if t.root == nil {
		t.root = &node{key: el}
		return
	}

	cur := t.root
	for cur != nil {
		if el < cur.key {
			if cur.left == nil {
				cur.left = &node{key: el}
				break
			}
			cur = cur.left
			continue
		} else {
			if cur.right == nil {
				cur.right = &node{key: el}
				break
			}
			cur = cur.right
			continue
		}
	}
}

func (t *sortTree) getAll() []int {
	return t.get(t.root)
}

func (t *sortTree) get(subroot *node) []int {
	if subroot == nil {
		return []int{}
	}
	var l, r []int
	if subroot.left != nil {
		l = t.get(subroot.left)
	}
	if subroot.right != nil {
		r = t.get(subroot.right)
	}
	res := append(l, subroot.key)
	res = append(res, r...)
	return res
}
