package main

func moveZeroes(nums []int) {
	insertIndex := 0
	count := 0
	for _, val := range nums {
		if val == 0 {
			count += 1
		} else {
			nums[insertIndex] = val
			insertIndex += 1
		}
	}
	for count > 0 {
		nums[insertIndex] = 0
		insertIndex += 1
		count--
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
