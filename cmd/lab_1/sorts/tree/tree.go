package tree

import "github.com/horockey/PSI/cmd/lab_1/sorts"

var _ sorts.SortAlgo = &tree{}

type tree struct{}

func New() *tree {
	return &tree{}
}

func (srt *tree) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *tree) String() string {
	return "tree"
}
