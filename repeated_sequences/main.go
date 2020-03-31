package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	var ret []string

	length := len(s)
	maxStrLen := 10
	dic := make(map[string]int)
	for i := 0; i <= length - maxStrLen; i++ {
		dic[s[i: i + maxStrLen]] += 1
	}

	for k, v := range dic {
		if v > 1 {
			ret = append(ret, k)
		}
	}

	return ret
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}