package merge_sorted_array

import "sort"

// Example: 1
//
//	Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//	Output: [1,2,2,3,5,6]
var Nums11 = []int{1, 2, 3, 0, 0, 0}
var M1 = 3
var Nums12 = []int{2, 5, 6}
var N1 = 3

// Example: 2
//
//	Input: nums1 = [1], m = 1, nums2 = [], n = 0
//	Output: [1]
var Nums21 = []int{1}
var M2 = 3
var Nums22 = []int{}
var N2 = 0

// Example: 3
//
//	Input: nums1 = [0], m = 0, nums2 = [1], n = 1
//	Output: [1]
var Nums31 = []int{0}
var M3 = 0
var Nums32 = []int{1}
var N3 = 1

func Algorithm(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

	sort.Ints(nums1)
}
