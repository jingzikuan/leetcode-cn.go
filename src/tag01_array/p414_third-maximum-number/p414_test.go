package p414_third_maximum_number

import "testing"

func TestThirdMax(t *testing.T) {
	nums := []int{1, 1, 2}
	t.Log(nums)
	max := thirdMax(nums)
	t.Log(max)
}
