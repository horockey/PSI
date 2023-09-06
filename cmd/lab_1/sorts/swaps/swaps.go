package swaps

import "github.com/horockey/PSI/cmd/lab_1/sorts"

var _ sorts.SortAlgo = &swaps{}

type swaps struct{}

func New() *swaps {
	return &swaps{}
}

func (srt *swaps) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *swaps) String() string {
	return "swaps"
}
