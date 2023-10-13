package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	items := []*item{
		{
			name:   "Слиток золота",
			weight: 3.0,
			cost:   100,
		},
		{
			name:   "Слиток серебра",
			weight: 5.5,
			cost:   100,
		},
		{
			name:   "Железная руда",
			weight: 10.3,
			cost:   50,
		},
		{
			name:   "Куб тессеракта",
			weight: 15.4,
			cost:   123.45,
		},
	}

	const backpackMaxWeight float64 = 18.5

	best := bests{}

	// Induction base
	best[0] = &bestEntry{
		cost:  0.0,
		items: []*item{},
	}

	for _, item := range items {
		for weight, entry := range best {
			if slices.Contains(entry.items, item) {
				continue
			}
			if weight+item.weight > backpackMaxWeight {
				continue
			}

			if oldBest, found := best[weight+item.weight]; !found || entry.cost+item.cost > oldBest.cost {
				best[weight+item.weight] = &bestEntry{
					weight: weight + item.weight,
					cost:   entry.cost + item.cost,
					items:  append(slices.Clone(entry.items), item),
				}
			}
		}
	}

	bestSlice := best.toSlice()
	sort.Slice(bestSlice, func(i, j int) bool {
		return bestSlice[i].cost < bestSlice[j].cost
	})

	fmt.Println(bestSlice[len(bestSlice)-1].String())
}
