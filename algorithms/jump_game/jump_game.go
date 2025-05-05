package jump_game

func Algorithm(nums []int) bool {
	len := len(nums) - 1
	if len == 0 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	goal := len
	for i := len - 1; i >= 0; i-- {
		jumps := nums[i]
		if jumps+i >= goal {
			goal = i
		}
	}

	return goal == 0
}

func Algorithm3(nums []int) bool {
	len := len(nums)
	if len == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	for i := range len {
		jumps := nums[i]
		hasJumps := false
		for j := range jumps {
			index := i + j + 1
			if index > len-1 {
				continue
			}

			jump := nums[index]
			if jump == 0 && index != len-1 {
				if j+1 == jumps && !hasJumps {
					if i-1 >= 0 && nums[i-1] > index {
						i++
						break
					}

					return false
				}
			} else {
				hasJumps = true
			}

			if jump+index >= len-1 {
				return true
			}
		}
	}

	return true
}

func Algorithm2(nums []int) bool {
	len := len(nums)
	nlen := len - 1

	if len == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	idx := 0
	for i := range len {
		jumps := nums[i]
		for j := jumps; j > 0; j-- {
			if i+j == nlen || j == nlen {
				return true
			}

			if i+j > nlen {
				continue
			}

			if nums[i+j] > 0 {
				if i+j > idx {
					idx = i + j
				}
			}
		}
		if idx == nlen {
			return true
		}
	}

	return idx == nlen
}

func Algorithm1(nums []int) bool {
	nlen := len(nums)
	if nlen == 1 {
		return true
	}

	return jump(nums, 0, nlen) == nlen-1
}

func jump(nums []int, idx int, nlen int) int {
	if idx == nlen {
		return idx
	}

	if idx > nlen {
		return 0
	}

	jumps := nums[idx]
	if jumps == 0 {
		return 0
	}

	gidx := idx
	for j := jumps; j > 0; j-- {
		jidx := jump(nums, idx+j, nlen)

		if jidx > gidx {
			gidx = jidx
		}
	}

	return gidx
}
