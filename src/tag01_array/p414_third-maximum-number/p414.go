package p414_third_maximum_number

func thirdMax(nums []int) int {
	//用map去重
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 0
	}
	newNums := []int{}
	for k := range m {
		newNums = append(newNums, k)
	}
	//计算元素个数
	maxSize := len(newNums)
	if len(newNums) < 3 {
		//如果元素个数小于3，直接计算最大值并返回
		res := newNums[0]
		for k := 1; k < maxSize; k++ {
			if res < newNums[k] {
				res = newNums[k]
			}
		}
		return res
	}

	//循环3次，每次去掉一个最大值，把第三次取得的最大值res返回
	res := 0
	//fmt.Println(newNums)
	for k := 0; k < 3; k++ {
		t := newNums[0]
		ind := 0
		for i := 0; i < len(newNums); i++ {
			if t <= newNums[i] {
				t = newNums[i]
				ind = i
				res = newNums[ind]
				//fmt.Println(k, ind, res)
			}
		}
		//fmt.Println("==", k, ind, res)
		//fmt.Println(k, ind, newNums, newNums[:ind])
		tmpNums := append([]int{}, newNums[:ind]...)
		if ind < len(newNums)-1 {
			tmpNums = append(tmpNums, newNums[ind+1:]...)
		}
		newNums = append([]int{}, tmpNums...)
		//fmt.Println(k, newNums)
	}

	return res
}
