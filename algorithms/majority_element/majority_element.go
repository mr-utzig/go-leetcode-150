package majority_element

// Example 1:
//
//	Input: nums = [3,2,3]
//	Output: 3

// Example 2:
//
//	Input: nums = [2,2,1,1,1,2,2]
//	Output: 2

func Algorithm(nums []int) int {
	l := len(nums)
	h := l / 2
	e := make(map[int]int)
	m := 0

	for i := range l {
		v := nums[i]
		e[v] += 1

		if e[v] > h {
			m = v
			break
		}
	}

	return m
}
