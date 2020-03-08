package main

import (
	"container/list"
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	if deck == nil {
		return nil
	}
	ret := []int{}
	if len(deck) == 0 {
		return ret
	}

	sort.Slice(deck, func(i, j int) bool {
		return deck[i] < deck[j]
	})

	dl := list.New()
	length := len(deck)

	for i := length - 1; i >= 0; i-- {
		ele := deck[i]

		if dl.Len() == 0 {
			dl.PushBack(ele)
		} else {
			tail := dl.Back()
			dl.Remove(tail)
			dl.PushFront(tail.Value)
			dl.PushFront(ele)
		}
	}

	for ite := dl.Front(); ite != nil; ite = ite.Next() {
		ret = append(ret, ite.Value.(int))
	}
	return ret
}

func main() {
	data := []int{17, 13, 11, 2, 3, 5, 7}

	fmt.Println(deckRevealedIncreasing(data))
}
