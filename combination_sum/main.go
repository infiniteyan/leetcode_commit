package main

import (
	"fmt"
	"sort"
)

func dfs(candidates []int, start int, end int, path []int, result *[][]int, target int) {
	for i := end; i >= start; i-- {
		left := target - candidates[i]
		path = append(path, candidates[i])
		if left == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			*result = append(*result, temp)
			path = path[0 : len(path) - 1]
			continue
		} else if left < 0 {
			path = path[0 : len(path) - 1]
			continue
		} else {
			dfs(candidates, start, i, path, result, left)
			path = path[0 : len(path) - 1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	len := len(candidates)

	var path []int
	var result [][]int
	dfs(candidates, 0, len - 1, path, &result, target)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	ret := combinationSum(candidates, 7)
	fmt.Println(ret)
}