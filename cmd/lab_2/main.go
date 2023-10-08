package main

import (
	"fmt"
	"log"

	"github.com/horockey/PSI/pkg/matrixes"
)

func main() {
	matrs := []*matrixes.Matrix[int]{
		matrixes.New([][]int{
			{1, 2, 3},
			{4, 5, 6},
		}),
		matrixes.New([][]int{
			{7, 8, 9, 10},
			{11, 12, 13, 14},
			{15, 16, 17, 18},
		}),
		matrixes.New([][]int{
			{10},
			{20},
			{30},
			{40},
		}),
	}

	optimSolution := optimOrder(matrs...)
	fmt.Println(optimSolution)

	ordered := make([]*matrixes.Matrix[int], 0, len(matrs))
	for _, idx := range optimSolution.path {
		ordered = append(ordered, matrs[idx])
	}

	res, err := matrixes.MultFixOrder(ordered...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result:\n%s", res)
}
