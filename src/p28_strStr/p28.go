package p28_strStr

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}
	if len(haystack) != 0 && len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := range haystack {
		if i+len(needle) > len(haystack) {
			return -1
		}
		s := haystack[i : i+len(needle)]
		if s == needle {
			return i
		}
	}

	return -1
}

/**
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	sublen := len(haystack) - len(needle)
	if sublen < 0 {
		return -1
	}
	for i := 0; i <= sublen; i++ {
		newstr := haystack[i : i+len(needle)]
		if newstr == needle {
			return i
		}
	}
	return -1
}
