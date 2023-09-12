package main

import (
	"fmt"
)

func GetTheNumberBits(x int) []int {
	bits := []int{}
	for x != 0 {
		bits = append(bits, x%10)
		x = x / 10
	}
	return bits
}

func compress(chars []byte) int {
	startC := chars[0]
	insertIndex := 0
	lengthOfS := 0
	length := 1
	for _, c := range chars[1:] {
		if c != startC {
			chars[insertIndex] = startC
			insertIndex, lengthOfS, startC = insertIndex+1, lengthOfS+1, c
			if length != 1 {
				numberBits := GetTheNumberBits(length)
				for i, n := range numberBits {
					chars[insertIndex+len(numberBits)-1-i] = byte('0' + n)
				}
				insertIndex, lengthOfS, length = insertIndex+len(numberBits), lengthOfS+len(numberBits), 1
			}
			startC = c
		} else {
			length += 1
		}
	}
	chars[insertIndex] = startC
	insertIndex, lengthOfS = insertIndex+1, lengthOfS+1
	fmt.Println(insertIndex)
	if length != 1 {
		numberBits := GetTheNumberBits(length)
		for i, n := range numberBits {
			chars[insertIndex+len(numberBits)-1-i] = byte(n + '0')
		}
		lengthOfS += len(numberBits)
	}
	fmt.Printf("%s\n", chars)
	return lengthOfS
}
func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
