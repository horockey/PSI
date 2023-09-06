package shell

import "github.com/horockey/PSI/cmd/lab_1/sorts"

var _ sorts.SortAlgo = &shell{}

type shell struct{}

func New() *shell {
	return &shell{}
}

func (srt *shell) Sort(arr []int) []int {
	// TODO
	return nil
}

func (srt *shell) String() string {
	return "shell"
}
