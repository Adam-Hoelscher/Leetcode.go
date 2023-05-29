package main

func LargestValuesDepthInFunction(root *TreeNode) []int {

	var ans []int

	var d int
	var rf func(_ *TreeNode)
	rf = func(node *TreeNode) {

		if node == nil {
			return
		}

		if d == len(ans) {
			ans = append(ans, node.Val)
		} else if ans[d] < node.Val {
			ans[d] = node.Val
		}

		d++

		if node.Left != nil {
			rf(node.Left)
		}

		if node.Right != nil {
			rf(node.Right)
		}

		d--
	}

	rf(root)
	return ans
}

var d int

func LargestValuesDepthInModule(root *TreeNode) []int {

	var ans []int

	var rf func(_ *TreeNode)
	rf = func(node *TreeNode) {

		if node == nil {
			return
		}

		if d == len(ans) {
			ans = append(ans, node.Val)
		} else if ans[d] < node.Val {
			ans[d] = node.Val
		}

		d++

		if node.Left != nil {
			rf(node.Left)
		}

		if node.Right != nil {
			rf(node.Right)
		}

		d--
	}

	rf(root)
	return ans
}

func LargestValuesDepthInRecursion(root *TreeNode) []int {

	var ans []int

	var rf func(_ *TreeNode, _ int)
	rf = func(node *TreeNode, d int) {

		if d > 50 {
			return
		}
		if node == nil {
			return
		}

		if d == len(ans) {
			ans = append(ans, node.Val)
		} else if ans[d] < node.Val {
			ans[d] = node.Val
		}

		if node.Left != nil {
			rf(node.Left, d+1)
		}

		if node.Right != nil {
			rf(node.Right, d+1)
		}

	}

	rf(root, 0)
	return ans
}
