package main

import "fmt"

/*
给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
*/
func kidsWithCandies(candies []int, extraCandies int) (result []bool) {
	max := candies[0]
	for _, val := range candies {
		if val > max {
			max = val
		}
	}
	for _, val := range candies {
		if val+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
