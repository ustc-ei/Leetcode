package main

import (
	"fmt"
	datastructure "leetcode/DataStructure"
)

/*
给你一个整数数组 nums，返回 数组 answer ,

其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积.

题目数据保证数组 nums 之中任意元素的全部前缀元素和后缀的乘积都在 32 位整数范围内。

请不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/
func productExceptSelf(nums []int) (result []int) {
	S := datastructure.NewSegmentTree(nums)
	for i := range nums {
		if i == 0 {
			result = append(result, S.Query(1, len(nums)-1))
		} else if i == len(nums)-1 {
			result = append(result, S.Query(0, len(nums)-2))
		} else {
			result = append(result, S.Query(0, i-1)*S.Query(i+1, len(nums)-1))
		}
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
