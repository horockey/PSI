package merge

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &merge{}

type merge struct{}

func New() *merge {
	return &merge{}
}

func (srt *merge) Sort(src []int) []int {
	arr := make([]int, len(src))
	copy(arr, src)

	if len(arr) <= 1 {
		return arr
	}

	left := srt.Sort(arr[:len(arr)/2])
	right := srt.Sort(arr[len(arr)/2:])

	li, ri := 0, 0
	for li < len(left) && ri < len(right) {
		var v int
		if left[li] < right[ri] {
			v = left[li]
			li++
		} else {
			v = right[ri]
			ri++
		}
		arr[li+ri-1] = v
	}

	for li < len(left) {
		arr[li+ri] = left[li]
		li++
	}
	for ri < len(right) {
		arr[li+ri] = right[ri]
		ri++
	}

	return arr
}

func (srt *merge) String() string {
	return "merge"
}
