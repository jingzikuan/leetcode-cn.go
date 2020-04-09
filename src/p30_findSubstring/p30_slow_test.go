package p30_findSubstring

import (
	"fmt"
	"testing"
)

func TestFindSubstringSlow(t *testing.T) {
	s := "aaaaaaaa"
	words := []string{"aa", "aa", "aa"}
	res := findSubstringSlow(s, words)
	fmt.Println("res=", res)
}
