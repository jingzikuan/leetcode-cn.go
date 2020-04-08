package p1_two_sum

func TwoSum(nums []int, target int) []int {
	res := []int{}
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func TwoSumMap(nums []int, target int) []int {
	res := []int{}
	tmpMap := make(map[int][]int)
	for i, v := range nums {
		if _, ok := tmpMap[v]; ok {
			tmpMap[v] = append(tmpMap[v], i)
		} else {
			tmpMap[v] = []int{i}
		}
	}
	for i, v := range nums {
		if arr, ok := tmpMap[target-v]; ok {
			if v == target-v {
				if len(arr) > 1 {
					return append(res, arr[0], arr[1])
				}
			} else {
				return append(res, i, arr[0])
			}
		}
	}
	return res
}
