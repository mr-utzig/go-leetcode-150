package remove_duplicates_from_sorted_array

// Example 1:
//
//	Input: nums = [1,1,2]
//	Output: 2, nums = [1,2,_]

// Example 2:
//
//	Input: nums = [0,0,1,1,1,2,2,3,3,4]
//	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

func Algorithm(nums []int) int {
	for i := 0; i < len(nums); {
		if len(nums) > i+1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
