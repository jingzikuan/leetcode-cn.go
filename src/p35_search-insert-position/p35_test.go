package p35_search_insert_position

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	target := 4
	nums := []int{1, 3, 5}
	fmt.Println(target)
	res := searchInsert(nums, target)
	fmt.Println(res)
}
