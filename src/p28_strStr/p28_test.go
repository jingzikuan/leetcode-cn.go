package p28_strStr

import "testing"

func TestStrStr(t *testing.T) {
	//test
	if strStr("abcdef", "bc") != 1 {
		t.FailNow()
	}
}
