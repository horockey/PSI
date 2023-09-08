package merge_test

import (
	"testing"

	"github.com/horockey/PSI/pkg/sorts/merge"
	"github.com/stretchr/testify/require"
)

func TestSort_Normal(t *testing.T) {
	arr := []int{54, 9, -8, -41, 10}
	expected := []int{-41, -8, 9, 10, 54}

	res := merge.New().Sort(arr)

	require.EqualValues(t, expected, res)
}

func TestSort_Empty(t *testing.T) {
	arr := []int{}
	expected := []int{}

	res := merge.New().Sort(arr)

	require.EqualValues(t, expected, res)
}

func TestSort_Nil(t *testing.T) {
	var arr []int = nil
	expected := []int{}

	res := merge.New().Sort(arr)

	require.EqualValues(t, expected, res)
}
