package p283_moveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 0, 1}
	t.Log(nums)
	moveZeroes(nums)
	t.Log(nums)
}
func TestMoveZeroes2(t *testing.T) {
	nums := []int{0, 1, 0, 2, 3, 0, 4, 5}
	t.Log(nums)
	moveZeroes(nums)
	t.Log(nums)
}
