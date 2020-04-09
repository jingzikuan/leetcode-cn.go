package p28_strStr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	//test
	if strStr("abcdef", "bc") != 1 {
		t.FailNow()
	}
}

func TestStrStr2(t *testing.T) {
	res := strStr2("a", "a")
	fmt.Println(res)
	if res != 0 {
		t.FailNow()
	}
}
