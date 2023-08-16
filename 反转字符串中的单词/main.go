package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) (result string) {
	splitWords := strings.Split(s, " ")
	for index := len(splitWords) - 1; index >= 0; index-- {
		if splitWords[index] != "" {
			result += splitWords[index] + " "
		}
	}
	return result[:len(result)-1]
}

func main() {
	s := "  hello world  "
	fmt.Println(len(reverseWords(s)))
}
