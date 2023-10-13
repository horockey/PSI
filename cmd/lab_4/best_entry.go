package main

import (
	"fmt"
)

type bestEntry struct {
	weight float64
	cost   float64
	items  []*item
}

func (e *bestEntry) String() string {
	res := "Items:\n"
	for idx, item := range e.items {
		res += fmt.Sprintf("%d. %s\n", idx+1, item.String())
	}
	res += "\n"
	res += fmt.Sprintf("Total cost: %.2f\n", e.cost)
	res += fmt.Sprintf("Total weight: %.2f\n", e.weight)
	return res
}
