package remove_element

// Example 1:
//
//	Input: nums = [3,2,2,3], val = 3
//	Output: 2, nums = [2,2,_,_]
var Nums1 = []int{3, 2, 2, 3}
var Val1 = 3

// Example 2:
//
//	Input: nums = [0,1,2,2,3,0,4,2], val = 2
//	Output: 5, nums = [0,1,4,0,3,_,_,_]
var Nums2 = []int{0, 1, 2, 2, 3, 0, 4, 2}
var Val2 = 2

func Algorithm(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if len(nums) > i && nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
