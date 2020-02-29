package main

import "fmt"

var dic = map[string]string{
	"1": "A",
	"2": "B",
	"3": "C",
	"4": "D",
	"5": "E",
	"6": "F",
	"7": "G",
	"8": "H",
	"9": "I",
	"10": "J",
	"11": "K",
	"12": "L",
	"13": "M",
	"14": "N",
	"15": "O",
	"16": "P",
	"17": "Q",
	"18": "R",
	"19": "S",
	"20": "T",
	"21": "U",
	"22": "V",
	"23": "W",
	"24": "X",
	"25": "Y",
	"26": "Z",
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	length := len(s)
	record := make([]int, length)

	s0 := string(s[0])
	if _, ok := dic[s0]; ok {
		record[0] = 1
	} else {
		record[0] = 0
	}

	for j := 1; j < length; j++ {
		charAt := string(s[j])
		_, ok := dic[charAt]
		if ok {
			record[j] = record[j - 1]
			recentTwo := s[j - 1 : j + 1]
			_, ok = dic[recentTwo]
			if ok {
				if j >= 2 {
					record[j] += record[j - 2]
				} else {
					record[j] += 1
				}
			}
		} else {
			recentTwo := s[j - 1 : j + 1]
			_, ok = dic[recentTwo]
			if ok {
				if j >= 2 {
					record[j] = record[j - 2]
				} else {
					record[j] = 1
				}
			} else {
				record[j] = 0
			}
		}
	}

	return record[length - 1]
}

func main() {
	str := "09"
	fmt.Println(numDecodings(str))
}