package p349_intersection_of_two_arrays

import "sort"

/**
 * 先使用map去重，迭代map1到map2查询，存在的数据即是交集
 */
func intersectionMap(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	for _, n := range nums1 {
		map1[n] = 1
	}
	map2 := make(map[int]int)
	for _, n := range nums2 {
		map2[n] = 1
	}
	res := []int{}
	for key := range map1 {
		if _, ok := map2[key]; ok {
			res = append(res, key)
		}
	}
	return res
}

/**
 * 网友方案: 排序+双指针+去重
 */
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	i, j := 0, 0
	var last = -1 << 31
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if last != nums1[i] {
				res = append(res, nums1[i])
				last = nums1[i]
			}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
