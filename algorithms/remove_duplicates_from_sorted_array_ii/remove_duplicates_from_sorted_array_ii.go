package remove_duplicates_from_sorted_array_ii

// Example 1:
//
//	Input: nums = [1,1,1,2,2,3]
//	Output: 5, nums = [1,1,2,2,3,_]

// Example 2:
//
//	Input: nums = [0,0,1,1,1,1,2,3,3]
//	Output: 7, nums = [0,0,1,1,2,3,3,_,_]

func Algorithm(nums []int) int {
	for i := 0; i < len(nums); {
		if len(nums) > i+2 && nums[i] == nums[i+1] && nums[i] == nums[i+2] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
