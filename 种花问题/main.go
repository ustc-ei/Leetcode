package main

import "fmt"

/*
假设有一个很长的花坛, 一部分地块种植了花, 另一部分却没有。

可是, 花不能种植在相邻的地块上, 它们会争夺水源, 两者都会死去。

给你一个整数数组 flowerbed 表示花坛, 由若干 0 和 1 组成,

其中 0 表示没种植花, 1 表示种植了花.

另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花?

能则返回 true ，不能则返回 false 。
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	flower := make([]bool, len(flowerbed))
	for index, val := range flowerbed {
		if val == 1 {
			flower[index] = true
			if index != 0 {
				flower[index-1] = true
			}
			if index != len(flowerbed)-1 {
				flower[index+1] = true
			}
		}
	}
	var candidateIndex []int
	for index, val := range flower {
		if !val {
			candidateIndex = append(candidateIndex, index)
		}
	}
	if len(candidateIndex) < n {
		return false
	} else if len(candidateIndex) == 1 && n == 1 {
		return true
	}
	// fmt.Println(candidateIndex)
	num := 0
	for start, end := 0, 1; end < len(candidateIndex); {
		if candidateIndex[end-1]+1 == candidateIndex[end] {
			end += 1
		} else {
			num += (end - start + 1) / 2
			start = end
			end += 1
		}
		fmt.Println(start, end)
		if end == len(candidateIndex) {
			num += (end - start + 1) / 2
		}
	}
	if num >= n {
		return true
	} else {
		return false
	}
}
func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
