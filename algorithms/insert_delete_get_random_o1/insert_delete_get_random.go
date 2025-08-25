package insertdeletegetrandomo1

import (
	"math/rand"
)

type RandomizedSet struct {
	Set  map[int]bool
	Size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Set: make(map[int]bool),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if rs.Set[val] {
		return false
	}

	rs.Set[val] = true
	rs.Size++

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if !rs.Set[val] {
		return false
	}

	delete(rs.Set, val)
	rs.Size--

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	keys := make([]int, 0, rs.Size)
	for k := range rs.Set {
		keys = append(keys, k)
	}

	k := rand.Intn(rs.Size)
	return keys[k]
}
