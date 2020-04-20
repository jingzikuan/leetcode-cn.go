package p349_intersection_of_two_arrays

import "testing"

func TestIntersection(t *testing.T) {
	var nums1 []int = []int{1, 1, 4, 5, 3, 2}
	var nums2 []int = []int{1, 1, 3, 7, 9}
	t.Log("nums1", nums1)
	t.Log("nums2", nums2)
	res := intersection(nums1, nums2)
	t.Log("nums1", nums1)
	t.Log("nums2", nums2)
	t.Log(res)
}
