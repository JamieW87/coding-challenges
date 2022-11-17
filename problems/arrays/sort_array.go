package problems

import (
	"sort"
)

//Sort an unsorted array and return the median (middle number)

func SortArray(a []int32) int32 {

	sort.Slice(a, func(i int, j int) bool {
		return a[i] < a[j]
	})

	div := len(a) / 2

	return a[div]
}
