package main

import (
	"fmt"
)

/*
对于字符串 s 和 t, 只有在 s = t + ... + t (t 自身连接 1 次或多次) 时, 我们才认定 "t" 能除尽 "s"。
给定两个字符串 str1 和 str2。返回最长字符串 x, 要求满足 x 能除尽 str1 且 x 能除尽 str2 。
*/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	length := gcd(len(str1), len(str2))
	result := ""
	result += str1[:length+1]
	return result
}

func main() {
	str1 := "ABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
