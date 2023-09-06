package merge

import "github.com/horockey/PSI/cmd/lab_1/sorts"

var _ sorts.SortAlgo = &merge{}

type merge struct{}

func New() *merge {
	return &merge{}
}

func (srt *merge) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *merge) String() string {
	return "merge"
}
