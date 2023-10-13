package main

type bests map[float64]*bestEntry

func (b bests) toSlice() []*bestEntry {
	res := make([]*bestEntry, 0, len(b))
	for _, entry := range b {
		res = append(res, entry)
	}
	return res
}
