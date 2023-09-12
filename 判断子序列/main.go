package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i, j := 0, 0
	flag := false
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				flag = true
				break
			}
		}
		j++
	}
	return flag
}

func main() {
	s, t := "acb", "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
