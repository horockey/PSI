package swaps

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &swaps{}

type swaps struct{}

func New() *swaps {
	return &swaps{}
}

func (srt *swaps) Sort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 0; i < len(res); i++ {
		for j := len(res) - 1; j > i; j-- {
			if res[j] < res[j-1] {
				res[j], res[j-1] = res[j-1], res[j]
			}
		}
	}

	return res
}

func (srt *swaps) String() string {
	return "swaps"
}
