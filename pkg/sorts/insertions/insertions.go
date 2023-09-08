package insertions

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &insertions{}

type insertions struct{}

func New() *insertions {
	return &insertions{}
}

func (srt *insertions) Sort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 1; i < len(res); i++ {
		el := res[i]
		for j := i; j > 0; j-- {
			if el < res[j-1] {
				res[j] = res[j-1]
				res[j-1] = el
			}
		}
	}

	return res
}

func (srt *insertions) String() string {
	return "insertions"
}
