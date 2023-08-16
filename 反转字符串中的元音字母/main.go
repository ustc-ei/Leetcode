package main

import "fmt"

/*
给你一个字符串 s, 仅反转字符串中的所有元音字母, 并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u', 且可能以大小写两种形式出现不止一次。
*/
func reverseVowels(s string) string {
	vowelMaps := make(map[rune]bool)
	reverseRunes := []rune{}
	vowelrunes := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, val := range vowelrunes {
		vowelMaps[val] = true
	}
	for _, val := range s {
		if ok := vowelMaps[val]; ok {
			reverseRunes = append(reverseRunes, val)
		}
	}
	lengthReverseRunes := len(reverseRunes)
	result := []rune{}
	for index, vowelNum := 0, 0; index < len(s); index++ {
		if ok := vowelMaps[rune(s[index])]; !ok {
			result = append(result, rune(s[index]))
		} else {
			result = append(result, rune(reverseRunes[lengthReverseRunes-1-vowelNum]))
			vowelNum += 1
		}
	}
	return string(result)
}

func main() {
	s := "hello"
	fmt.Println(reverseVowels(s))
}
