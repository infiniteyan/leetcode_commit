package main

import "fmt"

func in(val int, dataMap *map[int]int) bool {
	_, ok := (*dataMap)[val]
	return ok
}

func findMinHeightTrees(n int, edges [][]int) []int {
	var result []int

	for i := 0; i < n; i++ {
		result = append(result, i)
	}

	for ;len(result) > 2; {
		length := len(edges)
		var digitMap = map[int]int{}
		for i := 0; i < length; i++ {
			val := edges[i][0]
			if _, ok := digitMap[val]; ok {
				digitMap[val]++
			} else {
				digitMap[val] = 1
			}
			val = edges[i][1]
			if _, ok := digitMap[val]; ok {
				digitMap[val]++
			} else {
				digitMap[val] = 1
			}
		}

		leafMap := make(map[int]int, 0)
		for k, v := range digitMap {
			if v == 1 {
				leafMap[k] = 1
			}
		}

		tempEdges := make([][]int, 0)
		for i := 0; i < length; i++ {
			edge := edges[i]
			if in(edge[0], &leafMap) || in(edge[1], &leafMap) {
				continue
			}
			tempEdges = append(tempEdges, edge)
		}

		tempRet := make([]int, 0)
		for i := 0; i < len(result); i++ {
			if in(result[i], &leafMap) {
				continue
			}
			tempRet = append(tempRet, result[i])
		}
		edges = tempEdges
		result = tempRet
	}

	return result
}

func main() {
	n := 6
	edges := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{3, 4},
		[]int{4, 5},
	}

	fmt.Println(findMinHeightTrees(n, edges))
}
