package main

import "fmt"

func clumsy(N int) int {
	loops := N / 4
	left := N % 4

	var ret int
	for i := 0; i < loops; i++ {
		base := N - i * 4
		var elementValue int
		if i == 0 {
			elementValue = base * (base - 1) / (base - 2) + (base - 3)
			ret += elementValue
		} else {
			elementValue = base * (base - 1) / (base - 2) - (base - 3)
			ret -= elementValue
		}
	}

	if left > 0 {
		if (loops > 0) {
			switch left {
			case 1:
				ret -= left
			case 2:
				ret -= left * (left - 1)
			case 3:
				ret -= left * (left - 1) * (left - 2)
			}
		} else {
			switch left {
			case 1:
				ret += left
			case 2:
				ret += left * (left - 1)
			case 3:
				ret += left * (left - 1) * (left - 2)
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(clumsy(1))
	fmt.Println(clumsy(4))
	fmt.Println(clumsy(10))
}