package main

import "fmt"
import "container/list"

func isDoubleDot(list *list.List) bool {
	if list.Len() < 2 {
		return false
	}

	if list.Back().Value.(string) == "." {
		list.Remove(list.Back())
		if list.Back().Value.(string) == "." {
			list.PushBack(".")
			return true
		}
		list.PushBack(".")
	}

	return false
}

func simplifyPath(path string) string {
	length := len(path)

	if length == 0 {
		return ""
	}
	stack := list.New()
	for i := 0; i < length; i++ {
		if string(path[i]) == "/" && stack.Len() != 0 {
			if stack.Back().Value.(string) == "/" { //遇到连续的/则丢弃一个
				continue
			} else {
				if stack.Back().Value.(string) == "." {
					if isDoubleDot(stack) {
						stack.Remove(stack.Back())
						stack.Remove(stack.Back())
						stack.Remove(stack.Back())

						for stack.Len() != 0 && stack.Back().Value.(string) != "/" {
							stack.Remove(stack.Back())
						}

						if stack.Len() == 0 {
							stack.PushBack(string(path[i]))
						}

						continue
					} else {
						stack.Remove(stack.Back())
						stack.Remove(stack.Back())
					}
					stack.PushBack(string(path[i]))
				} else {
					stack.PushBack(string(path[i]))
				}

			}
		} else {
			stack.PushBack(string(path[i]))
		}
	}

	if stack.Back().Value.(string) == "/" && stack.Len() > 1 {
		stack.Remove(stack.Back())
	}

	if stack.Back().Value.(string) == "." {
		if isDoubleDot(stack) {
			stack.Remove(stack.Back())
			stack.Remove(stack.Back())
			stack.Remove(stack.Back())

			for stack.Len() != 0 && stack.Back().Value.(string) != "/" {
				stack.Remove(stack.Back())
			}

			if stack.Len() != 0 {
				stack.Remove(stack.Back())
			}
		} else {
			stack.Remove(stack.Back())
			if stack.Len() > 1 {
				stack.Remove(stack.Back())
			}
		}
	}

	var result string
	for stack.Len() != 0 {
		result += stack.Front().Value.(string)
		stack.Remove(stack.Front())
	}

	if len(result) == 0 {
		result = "/"
	}

	return result
}

func main() {
	s := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(s))
}