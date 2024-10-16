package leetcode

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	// Could optimize the search for the solution by using a binary search
	// However, this is not important since the sort is much more expensive
	for i := 0; i < len(citations); i++ {
		if citations[i] == i+1 {
			return citations[i]
		} else if citations[i] < i+1 {
			return i
		}
	}
	return len(citations)
}
