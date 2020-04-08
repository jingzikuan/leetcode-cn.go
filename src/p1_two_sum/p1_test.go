package p1_two_sum

import "testing"

func TestTwoSumMap(t *testing.T) {
	arr := []int{2, 7, 11, 15, 7, 15}
	sum := 9
	res := TwoSumMap(arr, sum)
	t.Log(res)
	if len(res) != 2 {
		t.Error("返回空")
		t.FailNow()
	}
	if arr[res[0]]+arr[res[1]] != sum {
		t.FailNow()
	}
}

func TestTwoSumMap2(t *testing.T) {
	arr := []int{3, 2, 4}
	sum := 6
	res := TwoSumMap(arr, sum)
	t.Log(res)
	if len(res) != 2 {
		t.Error("返回空")
		t.FailNow()
	}
	if arr[res[0]]+arr[res[1]] != sum {
		t.FailNow()
	}
}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	sum := 9
	res := TwoSum(arr, sum)
	t.Log(res)
	if len(res) != 2 {
		t.Error("返回空")
		t.FailNow()
	}
	if arr[res[0]]+arr[res[1]] != sum {
		t.FailNow()
	}
}

func TestTwoSum2(t *testing.T) {
	arr := []int{3, 2, 4}
	sum := 6
	res := TwoSum(arr, sum)
	t.Log(res)
	if len(res) != 2 {
		t.Error("返回空")
		t.FailNow()
	}
	if arr[res[0]]+arr[res[1]] != sum {
		t.FailNow()
	}
}
