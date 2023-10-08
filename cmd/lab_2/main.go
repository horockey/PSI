package main

import (
	"fmt"
	"log"

	"github.com/horockey/PSI/pkg/matrixes"
)

func main() {
	m1 := matrixes.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
	})

	m2 := matrixes.New([][]int{
		{7, 8, 9, 10},
		{11, 12, 13, 14},
		{15, 16, 17, 18},
	})

	m3 := matrixes.New([][]int{
		{10},
		{20},
		{30},
		{40},
	})

	res, err := matrixes.Mult(m1, m2, m3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
