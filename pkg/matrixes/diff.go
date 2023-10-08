package matrixes

import (
	"fmt"
)

func Diff[T Number](a, b *Matrix[T]) (*Matrix[T], error) {
	if a.rows != b.rows || a.cols != b.cols {
		return nil, fmt.Errorf("invalid sizes: (%dx%d) x (%dx%d)", a.rows, a.cols, b.rows, b.cols)
	}

	res := make([][]T, a.rows, a.cols)

	for i := 0; i < int(a.rows); i++ {
		for j := 0; j < int(a.cols); j++ {
			res[i][j] = a.storage[i][j] - b.storage[i][j]
		}
	}

	return &Matrix[T]{
		storage: res,
		rows:    a.rows,
		cols:    a.cols,
	}, nil
}
