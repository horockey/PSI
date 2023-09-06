package sorts

type SortAlgo interface {
	Sort(arr []int) []int
	String() string
}
