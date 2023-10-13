package main

import "fmt"

type item struct {
	name   string
	cost   float64
	weight float64
}

func (i *item) String() string {
	return fmt.Sprintf("%s (%.2frub/%.2fkg)", i.name, i.cost, i.weight)
}
