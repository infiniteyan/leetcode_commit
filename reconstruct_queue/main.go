package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	var result = make([][]int, 0)

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] <= people[j][1]
		}

		return people[i][0] > people[j][0]
	})


	for i := 0; i < len(people); i++ {
		length := len(result)

		counter := people[i][1]
		var j int
		for j = 0; j < length && counter > 0; j++ {
			counter--
		}

		leftHalf := result[0:j]
		rightHalf := result[j:]
		temp := append([][]int{}, leftHalf...)
		temp = append(temp, people[i])
		temp = append(temp, rightHalf...)
		result = temp
	}

	return result
}

func main() {
	people := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2}}

	fmt.Println(reconstructQueue(people))
}
