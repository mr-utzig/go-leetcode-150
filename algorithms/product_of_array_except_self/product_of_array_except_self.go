package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := range answer {
		answer[i] = 1
	}

	prefix := 1
	for i := range n {
		answer[i] *= prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
