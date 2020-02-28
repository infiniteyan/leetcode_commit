package main

import "fmt"

var dic = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func dfs(digits string, startIndex int, length int, base string, result *[]string) {
	if startIndex >= length {
		*result = append(*result, base)
		return
	}

	digit := string(digits[startIndex])
	repStr := dic[digit]

	repLen := len(repStr)
	for i := 0; i < repLen; i++ {
		base += string(repStr[i])
		dfs(digits, startIndex + 1, length, base, result)
		base = base[0 : len(base) - 1]
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var baseStr string
	var result []string
	dfs(digits, 0, len(digits), baseStr, &result)
	return result
}

func main() {
	ret := letterCombinations("2346")
	fmt.Println(ret)
}