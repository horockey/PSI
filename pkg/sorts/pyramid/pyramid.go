package pyramid

import "github.com/horockey/PSI/pkg/sorts"

var _ sorts.SortAlgo = &pyramid{}

type pyramid struct{}

func New() *pyramid {
	return &pyramid{}
}

func (srt *pyramid) Sort(src []int) []int {
	arr := make([]int, len(src))
	copy(arr, src)

	for i := len(arr)/2 - 1; i >= 0; i-- {
		srt.heapify(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		srt.heapify(arr, i, 0)
	}

	return arr
}

func (srt *pyramid) String() string {
	return "pyramid"
}

func (srt *pyramid) heapify(arr []int, len int, subrootIdx int) {
	gri := subrootIdx
	li := 2*subrootIdx + 1
	ri := 2*subrootIdx + 2

	if li < len && arr[gri] < arr[li] {
		gri = li
	}
	if ri < len && arr[gri] < arr[ri] {
		gri = ri
	}

	if gri != subrootIdx {
		arr[gri], arr[subrootIdx] = arr[subrootIdx], arr[gri]
		srt.heapify(arr, len, gri)
	}
}
