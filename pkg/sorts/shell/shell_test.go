package shell_test

import (
	"testing"

	"github.com/horockey/PSI/pkg/sorts/shell"
	"github.com/stretchr/testify/require"
)

func TestSort_Normal(t *testing.T) {
	arr := []int{54, 9, -8, -41, 10}
	expected := []int{54, 9, -8, -41, 10}

	res := shell.New().Sort(arr)

	require.EqualValues(t, expected, res)
}

func TestSort_Empty(t *testing.T) {
	arr := []int{}
	expected := []int{}

	res := shell.New().Sort(arr)

	require.EqualValues(t, expected, res)
}

func TestSort_Nil(t *testing.T) {
	var arr []int = nil
	expected := []int{}

	res := shell.New().Sort(arr)

	require.EqualValues(t, expected, res)
}
