package matrixes

import (
	"fmt"
)

func mult2[T Number](a, b *Matrix[T]) (*Matrix[T], error) {
	if a.cols != b.rows {
		return a, fmt.Errorf("invalid sizes: (%dx%d) x (%dx%d)", a.rows, a.cols, b.rows, b.cols)
	}

	res := make([][]T, a.rows)
	for i := 0; i < len(res); i++ {
		res[i] = make([]T, b.cols)
	}

	for i := 0; i < int(a.rows); i++ {
		for j := 0; j < int(b.cols); j++ {
			for k := 0; k < int(a.cols); k++ {
				res[i][j] += a.storage[i][k] * b.storage[k][j]
			}
		}
	}

	return &Matrix[T]{
		storage: res,
		rows:    a.rows,
		cols:    b.cols,
	}, nil
}

// Multiplies given matrs in given order
func Mult[T Number](matrs ...*Matrix[T]) (*Matrix[T], error) {
	if len(matrs) == 0 {
		return &Matrix[T]{}, nil
	}
	if len(matrs) == 1 {
		return matrs[0], nil
	}

	last := matrs[0]
	for idx := 1; idx < len(matrs); idx++ {
		var err error
		last, err = mult2(last, matrs[idx])
		if err != nil {
			return nil, fmt.Errorf("multiplying on pos %d: %w", idx, err)
		}
	}

	return last, nil
}

// Multiplies given matrs. If sizes doesnt match, changes places
func MultFixOrder[T Number](matrs ...*Matrix[T]) (*Matrix[T], error) {
	if len(matrs) == 0 {
		return &Matrix[T]{}, nil
	}
	if len(matrs) == 1 {
		return matrs[0], nil
	}

	last := matrs[0]
	for idx := 1; idx < len(matrs); idx++ {
		var err error
		last, err = mult2(last, matrs[idx])
		if err != nil {
			origErr := err
			last, err = mult2(matrs[idx], last)
			if err != nil {
				return nil, fmt.Errorf("multiplying on pos %d: %w", idx, origErr)
			}
		}
	}

	return last, nil
}
