package quick

import (
	"math/rand"
	"time"

	"github.com/horockey/PSI/cmd/lab_1/sorts"
)

var _ sorts.SortAlgo = &quick{}

type quick struct {
	r rand.Rand
}

func New() *quick {
	return &quick{
		r: *rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (srt *quick) Sort(src []int) []int {
	arr := make([]int, len(src))
	copy(arr, src)
	srt.sort(arr, 0, len(arr)-1)
	return arr
}

func (srt *quick) String() string {
	return "quick"
}

func (srt *quick) sort(arr []int, l, r int) {
	if l >= r {
		return
	}

	border := srt.order(arr, l, r)
	srt.sort(arr, l, border-1)
	srt.sort(arr, border+1, r)
}

func (str *quick) order(arr []int, l, r int) int {
	// anchorIdx := str.r.Intn(len(arr))
	anchorIdx := r
	anchor := arr[anchorIdx]
	k := l
	for i := l; i < r; i++ {
		if arr[i] < anchor {
			arr[k], arr[i] = arr[i], arr[k]
			k++
		}
	}
	arr[anchorIdx], arr[k] = arr[k], arr[anchorIdx]
	return k
}
