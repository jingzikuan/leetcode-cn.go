package p283_moveZeroes

import "fmt"

//我的解
func moveZeroes2(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for k := i; k < len(nums)-1; k++ {
				nums[k] = nums[k+1]
				nums[k+1] = 0
			}
		}
	}
}

//最优解
func moveZeroes(nums []int) {
	j := 0
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}

			j++
		}
		fmt.Println(nums)
	}
}
