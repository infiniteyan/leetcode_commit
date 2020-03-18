package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type NestedInteger struct {
	val 	int
	childs 	[]*NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return n.val == 0
}

func (n NestedInteger) GetInteger() int {
	return n.val
}

func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

func (n *NestedInteger) Add(elem NestedInteger) {
	n.childs = append(n.childs, &elem)
}

func (n NestedInteger) GetList() []*NestedInteger {
	return n.childs
}

func reverseStr(s string) string {
	var result string
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

func convertToValue(s string, val *int) bool {
	if len(s) == 0 {
		return false
	}

	s = strings.Trim(s, "[]")
	value, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	*val = value
	return true
}

func deserialize(s string) *NestedInteger {
	result := &NestedInteger{}
	if len(s) == 0 {
		return nil
	}

	stack := list.New()
	length := len(s)

	var child []*NestedInteger
	for i := 0; i < length; i++ {
		if s[i] != ']' {
			stack.PushBack(string(s[i]))
		} else {
			var str string
			for stack.Len() != 0 {
				back := stack.Back()
				stack.Remove(stack.Back())
				if back.Value.(string) != "[" {
					str += back.Value.(string)
				} else {
					break
				}
			}

			if stack.Len() != 0 && stack.Back().Value.(string) == "," {
				stack.Remove(stack.Back())
			}
			value := 0
			success := convertToValue(reverseStr(str), &value)
			if child == nil {
				item := NestedInteger{}
				if success {
					item.SetInteger(value)
				}
				child = append([]*NestedInteger{}, &item)
			} else {
				item := NestedInteger{}
				if success {
					item.SetInteger(value)
				}
				for _, v := range child {
					item.Add(*v)
				}
				child = append([]*NestedInteger{}, &item)
			}
		}
	}

	if stack.Len() != 0 {
		var str string
		for stack.Len() != 0 {
			str += stack.Front().Value.(string)
			stack.Remove(stack.Front())
		}

		var value int
		if convertToValue(str, &value) {
			result.SetInteger(value)
		}
	} else {
		result = child[0]
	}

	return result
}

func main() {
	s := "[123,[456,[789]]]"
	nestedInteger := deserialize(s)
	fmt.Println(nestedInteger.val)
}