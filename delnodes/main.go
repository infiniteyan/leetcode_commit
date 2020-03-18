package main

import "fmt"

type TreeNode struct {
	Val 		int
	Left 		*TreeNode
	Right 		*TreeNode
}

func exist(slice []int, e int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == e {
			return true
		}
	}

	return false
}

func dfs(root *TreeNode, delete []int, ret *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var left *TreeNode = nil
	var right *TreeNode = nil
	if root.Left != nil {
		left = dfs(root.Left, delete, ret)
	}
	if root.Right != nil {
		right = dfs(root.Right, delete, ret)
	}

	if exist(delete, root.Val) {
		if left != nil {
			*ret = append(*ret, left)
		}
		if right != nil {
			*ret = append(*ret, right)
		}
		return nil
	} else {
		root.Left = left
		root.Right = right
		return root
	}
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var ret = []*TreeNode{}
	if root == nil {
		return ret
	}

	ret = append(ret, root)
	dfs(root, to_delete, &ret)
	return ret
}

func main() {
	one := TreeNode{Val:1, Left:nil, Right:nil}
	two := TreeNode{Val:2, Left:nil, Right:nil}
	three := TreeNode{Val:3, Left:nil, Right:nil}
	four := TreeNode{Val:4, Left:nil, Right:nil}
	five := TreeNode{Val:5, Left:nil, Right:nil}
	six := TreeNode{Val:6, Left:nil, Right:nil}
	seven := TreeNode{Val:7, Left:nil, Right:nil}
	two.Left = &four
	two.Right = &five
	three.Left = &six
	three.Right = &seven
	one.Left = &two
	one.Right = &three

	to_delete := []int{3, 5}

	allTrees := delNodes(&one, to_delete)
	fmt.Println(len(allTrees))
}
