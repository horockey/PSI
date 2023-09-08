package tree

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &tree{}

type tree struct {
	t *sortTree
}

func New() *tree {
	return &tree{
		t: &sortTree{},
	}
}

func (srt *tree) Sort(arr []int) []int {
	for _, el := range arr {
		srt.t.insert(el)
	}

	return srt.t.getAll()
}

func (srt *tree) String() string {
	return "tree"
}
