package hindex

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	s := 0
	for i := range len(citations) {
		if citations[i] >= i+1 {
			s++
		}
	}

	return s
}
