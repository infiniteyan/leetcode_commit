package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

func convertToList(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}

	var head *TreeNode
	var tail *TreeNode
	head = root
	if root.Left == nil && root.Right == nil {
		tail = root
		return head, tail
	}

	leftPtr := root.Left
	rightPtr := root.Right
	if leftPtr != nil {
		root.Left = nil
		leftH, leftT := convertToList(leftPtr)
		root.Right = leftH
		if rightPtr != nil {
			rightH, rightT := convertToList(rightPtr)
			leftT.Right = rightH
			return head, rightT
		} else {
			return head, leftT
		}
	} else {
		_, rightT := convertToList(rightPtr)
		return head, rightT
	}
}

func flatten(root *TreeNode)  {
	convertToList(root)

	ptr := root
	for {
		if ptr == nil {
			break
		}

		fmt.Printf("%d-->", ptr.Val)
		ptr = ptr.Right
	}
	fmt.Println()
}


func main() {
	one := TreeNode{Val: 1, Left:nil, Right:nil}
	two := TreeNode{Val: 2, Left: nil, Right: nil}
	three := TreeNode{Val: 3, Left: nil, Right: nil}
	four := TreeNode{Val: 4, Left: nil, Right: nil}
	five := TreeNode{Val: 5, Left: nil, Right: nil}
	six := TreeNode{Val: 6, Left: nil, Right: nil}
	two.Left = &three
	two.Right = &four
	five.Right = &six
	one.Left = &two
	one.Right = &five
	flatten(&one)
}