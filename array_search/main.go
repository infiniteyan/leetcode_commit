package array_search

import "fmt"

func search(nums []int, target int) int {
	l := 0
	h := len(nums) - 1

	var leftHalf = false
	for ;l < h; {
		midIndex := l + (h - l) / 2
		mid := nums[midIndex]

		if mid == target {
			return midIndex
		}
		if nums[l] == target {
			return l
		}
		if nums[h] == target {
			return h
		}

		if mid > nums[l] {
			leftHalf = true
		} else {
			leftHalf = false
		}

		if leftHalf {
			if mid > target {
				if nums[l] > target {
					l = midIndex
				} else {
					h = midIndex
				}
			} else {
				l = midIndex
			}
		} else {
			if mid > target {
				h = midIndex
			} else {
				if nums[h] > target {
					l = midIndex
				} else {
					h = midIndex
				}
			}
		}
	}

	return -1;
}

func main() {
	var target int = 1
	var nums []int = []int{4, 5, 6, 7, 0, 1, 2, 9}

	fmt.Println(search(nums, target))
}