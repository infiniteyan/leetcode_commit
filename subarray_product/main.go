package main

import "fmt"

func dfs(nums []int, currentValue *int, index int, k int, counter *int) {
	if index >= len(nums) {
		return
	}

	*currentValue = *currentValue * nums[index]

	if *currentValue >= k {
		return
	}

	*counter += 1

	dfs(nums, currentValue, index + 1, k, counter)
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var ret int
	length := len(nums)
	for i := 0; i < length; i++ {
		value := 1
		counter := 0
		dfs(nums, &value, i, k, &counter)
		ret += counter
	}

	return ret
}

func main() {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100))
}