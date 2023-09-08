package pyramid

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &pyramid{}

type pyramid struct{}

func New() *pyramid {
	return &pyramid{}
}

func (srt *pyramid) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *pyramid) String() string {
	return "pyramid"
}
