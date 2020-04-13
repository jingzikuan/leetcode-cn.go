package p35_search_insert_position

func searchInsert(nums []int, target int) int {
	for i := 0; i < (len(nums)+1)/2; i++ {
		if nums[i] >= target {
			return i
		}
		if i+1 < len(nums) && nums[i] < target && nums[i+1] > target {
			return i + 1
		}
		k := len(nums) - 1 - i
		if nums[k] == target {
			return k
		}
		if nums[k] < target {
			return k + 1
		}
		if k-1 > 0 && nums[k-1] < target && nums[k] > target {
			return k
		}
	}
	return 0
}

//高手的代码
func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
