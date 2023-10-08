package main

import (
	"fmt"
	"sort"

	"github.com/horockey/PSI/pkg/matrixes"
)

type solution struct {
	operations uint
	path       []int
	pathString string
}

func (s *solution) String() string {
	return fmt.Sprintf("Operations: %d\nSolution: %s\nOrder: %+v", s.operations, s.pathString, s.path)
}

func optimOrder[T matrixes.Number](matrs ...*matrixes.Matrix[T]) *solution {
	sizes := make([]uint, 0, len(matrs))
	for _, matr := range matrs {
		sizes = append(sizes, matr.Rows())
	}
	sizes = append(sizes, matrs[len(matrs)-1].Cols())

	n := len(sizes) - 1
	best := make([][]solution, n)
	for i := range best {
		best[i] = make([]solution, n)
	}

	for i := range best {
		best[i][i] = solution{
			operations: 0,
			pathString: fmt.Sprint(i),
		}
	}

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			best[i][j] = solution{
				operations: ^uint(0),
				pathString: "",
			} // MAX

			for k := i; k < j; k++ {
				ops := best[i][k].operations +
					best[k+1][j].operations +
					sizes[i]*sizes[k+1]*sizes[j+1]
				if ops < best[i][j].operations {
					best[i][j].operations = ops
					best[i][j].pathString = fmt.Sprintf("(%s) * (%s)", best[i][k].pathString, best[k+1][j].pathString)
				}
			}
		}
	}

	bestSolution := &best[0][n-1]

	recoverPath(bestSolution)
	return bestSolution
}

func recoverPath(s *solution) {
	opened := 0
	operandsNesting := []struct {
		operand int
		nesting int
	}{}
	for _, c := range s.pathString {
		if c == '(' {
			opened++
		}
		if c == ')' {
			opened--
		}

		if c >= '0' && c <= '9' {
			var operand int = int(c - '0')
			operandsNesting = append(operandsNesting, struct {
				operand int
				nesting int
			}{
				operand: operand,
				nesting: opened,
			})
		}
	}

	sort.Slice(operandsNesting, func(i, j int) bool {
		return operandsNesting[i].nesting > operandsNesting[j].nesting
	})

	s.path = make([]int, 0, len(operandsNesting))
	for _, opn := range operandsNesting {
		s.path = append(s.path, opn.operand)
	}
}
