package findLargestValueInEachTreeRow

import . "leetcode/leetcodeStruct"

func LargestValues(root *TreeNode) []int {

	var ans []int

	var rf func(node *TreeNode, d int)
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
