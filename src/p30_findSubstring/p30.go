package p30_findSubstring

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	fmt.Println("s=", s, "words=", words)
	res := []int{}
	wordsNum := len(words)
	if s == "" || wordsNum == 0 {
		return res
	}
	wordLen := len(words[0])
	if wordsNum == 1 && len(s) == wordLen && s == words[0] {
		return append(res, 0)
	}
	wordsMap := make(map[string]int)
	for _, v := range words {
		if num, ok := wordsMap[v]; ok {
			num++
			wordsMap[v] = num
		} else {
			wordsMap[v] = 1
		}
	}
	fmt.Println(wordsMap)
	for i := 0; i+wordLen < len(s); i++ {
		headStr := s[i : i+wordLen]
		if _, ok := wordsMap[headStr]; ok {
			endIdx := i + wordsNum*wordLen
			if endIdx > len(s) {
				continue
			}
			wholeStr := s[i:endIdx]
			fmt.Println(i, headStr, wholeStr)

			//把wholeStr按固定word长度转换成map，然后跟wordsMap比较后相同的放入成功匹配
			tmpMap := make(map[string]int)
			for j := 0; j < len(wholeStr); j = j + wordLen {
				tmpStr := wholeStr[j : j+wordLen]
				if num, ok := tmpMap[tmpStr]; ok {
					num++
					tmpMap[tmpStr] = num
				} else {
					tmpMap[tmpStr] = 1
				}
			}

			//比较两个map
			flag := true
			for k := range wordsMap {
				if wordsMap[k] != tmpMap[k] {
					flag = false
				}
			}
			if flag {
				res = append(res, i)
			}
		}
	}
	return res
}
