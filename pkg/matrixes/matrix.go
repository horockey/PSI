package matrixes

import "fmt"

type Matrix[T Number] struct {
	rows    uint
	cols    uint
	storage [][]T
}

func New[T Number](data [][]T) *Matrix[T] {
	if len(data) == 0 {
		return &Matrix[T]{}
	}

	return &Matrix[T]{
		storage: data,
		rows:    uint(len(data)),
		cols:    uint(len(data[0])),
	}
}

func (m *Matrix[T]) Get(row, col uint) (T, error) {
	if row >= m.rows || col >= m.cols {
		return *new(T), fmt.Errorf("given index out of range: (%d, %d)", row, col)
	}

	return m.storage[row][col], nil
}

func (m *Matrix[T]) Set(row, col uint, val T) error {
	if row >= m.rows || col >= m.cols {
		return fmt.Errorf("given index out of range: (%d, %d)", row, col)
	}

	m.storage[row][col] = val
	return nil
}

func (m *Matrix[T]) Rows() uint {
	return m.rows
}

func (m *Matrix[T]) Cols() uint {
	return m.cols
}

func (m *Matrix[T]) String() string {
	res := ""
	for row := 0; row < int(m.Rows()); row++ {
		for col := 0; col < int(m.Cols()); col++ {
			res += fmt.Sprint(m.storage[row][col])
			if col < int(m.Cols())-1 {
				res += ", "
			}
		}
		res += "\n"
	}

	return res
}
