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
