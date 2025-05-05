package jump_game

import "testing"

func TestNums1(t *testing.T) {
	// Example 1:
	// Input: nums = [2,3,1,1,4]
	// Output: true
	nums1 := []int{2, 3, 1, 1, 4}
	res := Algorithm(nums1)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums1, res)
	}
}

func TestNums2(t *testing.T) {
	// Example 2:
	// Input: nums = [3,2,1,0,4]
	// Output: false
	nums2 := []int{3, 2, 1, 0, 4}
	res := Algorithm(nums2)

	if res != false {
		t.Errorf(`Algorithm(%v) = %v, expected false`, nums2, res)
	}
}

func TestNums3(t *testing.T) {
	nums3 := []int{1, 0, 1, 0}
	res := Algorithm(nums3)

	if res != false {
		t.Errorf(`Algorithm(%v) = %v, expected false`, nums3, res)
	}
}

func TestNums4(t *testing.T) {
	nums4 := []int{1, 2}
	res := Algorithm(nums4)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums4, res)
	}
}

func TestNums5(t *testing.T) {
	nums5 := []int{2, 0}
	res := Algorithm(nums5)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums5, res)
	}
}

func TestNums6(t *testing.T) {
	nums6 := []int{1, 2, 3}
	res := Algorithm(nums6)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums6, res)
	}
}

func TestNums7(t *testing.T) {
	nums7 := []int{1, 1, 1, 0}
	res := Algorithm(nums7)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums7, res)
	}
}

func TestNums8(t *testing.T) {
	nums8 := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	res := Algorithm(nums8)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums8, res)
	}
}

func TestNums9(t *testing.T) {
	nums9 := []int{2, 0, 0, 0, 2, 0, 0, 0}
	res := Algorithm(nums9)

	if res != false {
		t.Errorf(`Algorithm(%v) = %v, expected false`, nums9, res)
	}
}

func TestNums10(t *testing.T) {
	nums10 := []int{4, 1, 0, 2, 2, 4, 4, 4, 1, 2}
	res := Algorithm(nums10)

	if res != true {
		t.Errorf(`Algorithm(%v) = %v, expected true`, nums10, res)
	}
}
