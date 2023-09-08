package selections

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &selections{}

type selections struct{}

func New() *selections {
	return &selections{}
}

func (srt *selections) Sort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 0; i < len(res); i++ {
		min := res[i]
		minIdx := i
		for j := i; j < len(res); j++ {
			if res[j] < min {
				min = res[j]
				minIdx = j
			}
		}
		res[i], res[minIdx] = res[minIdx], res[i]
	}
	return res
}

func (srt *selections) String() string {
	return "selections"
}
