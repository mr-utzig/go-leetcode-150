package rotate_array

// Example 1:
//
//	Input: nums = [1,2,3,4,5,6,7], k = 3
//	Output:       [5,6,7,1,2,3,4]
var Nums1 = []int{1, 2, 3, 4, 5, 6, 7}
var K1 = 3

// Example 2:
//
//	Input: nums = [-1,-100,3,99], k = 2
//	Output: 	  [3,99,-1,-100]
var Nums2 = []int{-1, -100, 3, 99}
var K2 = 2

func Algorithm(nums []int, k int) {
	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := range nums {
		nums[i] = result[i]
	}
}
