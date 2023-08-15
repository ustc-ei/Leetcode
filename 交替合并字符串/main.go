package main

import (
	"fmt"
)

/*
给你两个字符串 word1 和 word2. 请你从 word1 开始, 通过交替添加的字母来合并字符串.
如果一个字符串比另外一个字符串长, 就将多出来的字母追加到合并字符串的末尾
* word1 字符串1
* word2 字符串2
* return 合并后的字符串
*/
func mergeAlternately(word1 string, word2 string) string {
	var mergeStr []rune
	loopEnd := len(word1)
	if loopEnd >= len(word2) {
		loopEnd = len(word2)
	}
	for i := 0; i < loopEnd; i++ {
		mergeStr = append(mergeStr, rune(word1[i]))
		mergeStr = append(mergeStr, rune(word2[i]))
	}
	if loopEnd < len(word1) {
		mergeStr = append(mergeStr, []rune(word1[loopEnd:])...)
	}
	if loopEnd < len(word2) {
		mergeStr = append(mergeStr, []rune(word2[loopEnd:])...)
	}
	return string(mergeStr)
}

func main() {
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
