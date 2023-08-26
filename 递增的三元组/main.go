package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k]

返回 true ；否则，返回 false 。
*/
/*
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	length := len(nums)
	minValue, maxValue := make([]int, length), make([]int, length)
	minValue[0], maxValue[length-1] = nums[0], nums[length-1]
	for i := 1; i < length; i++ {
		minValue[i] = min(nums[i], minValue[i-1])
	}
	for i := length - 2; i >= 0; i-- {
		maxValue[i] = max(nums[i], maxValue[i+1])
	}
	for i := length - 1; i >= 0; i-- {
		if maxValue[i] > nums[i] && minValue[i] < nums[i] {
			return true
		}
	}
	return false
}
*/

func increasingTriplet(nums []int) bool {
	firstValue, secondValue := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if firstValue >= n {
			firstValue = n
		} else if secondValue >= n {
			secondValue = n
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))
}
