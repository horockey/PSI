package shell

import (
	"github.com/horockey/PSI/pkg/sorts"
)

var _ sorts.SortAlgo = &shell{}

type shell struct{}

func New() *shell {
	return &shell{}
}

func (srt *shell) Sort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for _, step := range srt.generateSteps(len(res)) {
		idxs := []int{}
		for i := 0; i < step; i++ {
			for k := 0; i+k*step < len(res); k++ {
				idxs = append(idxs, i+k*step)
			}
		}

		for _, idx := range idxs {
			el := res[idx]
			for i := len(idxs) - 1; idxs[i] > idx; i-- {
				if el < res[idxs[i]-1] {
					res[idxs[i-1]] = res[idxs[i]]
					res[idxs[i]] = el
				}
			}
		}
	}

	return res
}

func (srt *shell) String() string {
	return "shell"
}

func (srt *shell) generateSteps(n int) []int {
	res := []int{}

	for i := n / 2; i >= 1; i /= 2 {
		res = append(res, i)
	}

	return res
}
