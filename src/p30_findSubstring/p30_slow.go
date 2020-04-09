package p30_findSubstring

import (
	"fmt"
	"strings"
)

func findSubstringCur(s string, words []string) bool {
	fmt.Println("findSubstringCur s=", s, "words=", words)
	for i, word := range words {
		idx := strings.Index(s, word)
		if idx == 0 {
			if len(words) == 1 {
				fmt.Println("findSubstringCur s=", s, "words=", words, "-->", true)
				return true
			}
			if len(words) > 1 {
				newWords := append([]string{}, words[:i]...)
				newWords = append(newWords, words[i+1:]...)
				newS := s[idx+len(word):]
				b := findSubstringCur(newS, newWords)
				fmt.Println("findSubstringCur s=", s, "words=", words, "-->", b)
				return b
			}
		}
	}
	fmt.Println("findSubstringCur s=", s, "words=", words, "-->", false)
	return false
}

func isExist(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func findSubstringSlow(s string, words []string) []int {
	fmt.Println("s=", s, "words=", words)
	res := []int{}
	for i, word := range words {
		idx := strings.Index(s, word)
		for idx > -1 && idx < len(s) {
			fmt.Println("loop i=", i, "word="+word+", idx=", idx)
			if len(words) == 1 {
				fmt.Println("------------>found ", idx)
				if isExist(res, idx) == false {
					res = append(res, idx)
				}
			}
			if len(words) > 1 {
				newWords := append([]string{}, words[:i]...)
				newWords = append(newWords, words[i+1:]...)
				fmt.Println("loop i=", i, ", newWords=", newWords, ", words=", words)
				newS := s[idx+len(word):]
				if findSubstringCur(newS, newWords) {
					fmt.Println("------------>found ", idx)
					if isExist(res, idx) == false {
						res = append(res, idx)
					}
				}
			}
			newIdx := strings.Index(s[idx+1:], word)
			if newIdx == -1 {
				break
			}
			idx = idx + 1 + newIdx
		}
	}
	return res
}
