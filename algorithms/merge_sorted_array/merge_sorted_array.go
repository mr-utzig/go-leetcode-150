package merge_sorted_array

import "sort"

// Example: 1
//
//	Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//	Output: [1,2,2,3,5,6]

// Example: 2
//
//	Input: nums1 = [1], m = 1, nums2 = [], n = 0
//	Output: [1]

// Example: 3
//
//	Input: nums1 = [0], m = 0, nums2 = [1], n = 1
//	Output: [1]

func Algorithm(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

	sort.Ints(nums1)
}
