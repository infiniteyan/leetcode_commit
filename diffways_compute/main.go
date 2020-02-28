package main

import (
	"fmt"
	"strconv"
)

func isDigit(s string) bool {
	if s == "+" || s == "-" || s == "*" {
		return false
	} else {
		return true
	}
}

func calValue(v1 int, v2 int, op string) int {
	var ret int
	if op == "+" {
		ret= v1 + v2
	} else if op == "-" {
		ret = v1 - v2
	} else if op == "*" {
		ret = v1 * v2
	}
	return ret
}

func diffWaysToCompute(input string) []int {
	ret := make([]int, 0)

	if len(input) == 0 {
		return ret
	}

	if len(input) == 1 {
		val, _ := strconv.Atoi(input)
		ret = append(ret, val)
		return ret
	}

	length := len(input)
	for i := 0; i < length; i++ {
		if !isDigit(string(input[i])) {
			strLeft := input[0:i]
			strRight := input[i+1:]
			left := diffWaysToCompute(strLeft)
			right := diffWaysToCompute(strRight)
			for _, v1 := range left {
				for _, v2 := range right {
					ret = append(ret, calValue(v1, v2, string(input[i])))
				}
			}
		}
	}

	return ret
}

func main() {
	input := "2*3-4*5"
	ret := diffWaysToCompute(input)
	fmt.Println(ret)
}
