package quick

import "github.com/horockey/PSI/cmd/lab_1/sorts"

var _ sorts.SortAlgo = &quick{}

type quick struct{}

func New() *quick {
	return &quick{}
}

func (srt *quick) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *quick) String() string {
	return "quick"
}
