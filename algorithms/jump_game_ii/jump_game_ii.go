package jump_game_ii

func Algorithm(nums []int) int {
	jumps := 0
	end := 0
	farthest := 0
	for i := range len(nums) - 1 {
		farthest = max(farthest, i+nums[i])
		if i == end {
			jumps++
			end = farthest
		}
	}

	return jumps
}
