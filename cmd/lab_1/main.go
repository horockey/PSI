package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/horockey/PSI/cmd/lab_1/sorts"
	"github.com/horockey/PSI/cmd/lab_1/sorts/insertions"
	"github.com/horockey/PSI/cmd/lab_1/sorts/selections"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Timestamp().Logger()

	algos := []sorts.SortAlgo{
		insertions.New(),
		selections.New(),
		// swaps.New(),
		// quick.New(),
		// tree.New(),
		// pyramid.New(),
		// shell.New(),
		// merge.New(),
	}

	sizes := []int{
		5_000,
		10_000,
		100_000,
		150_000,
	}

	for _, size := range sizes {
		logger.Info().
			Int("size", size).
			Msg("New size iteration started!")
		arr := generateRandom(size)
		for _, algo := range algos {
			ts := time.Now()
			algo.Sort(arr)
			dur := time.Since(ts)
			logger.Info().
				Str("algo", algo.String()).
				Str("dur", dur.String()).
				Send()
		}
	}
}

func generateRandom(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, r.Int())
	}
	return res
}
