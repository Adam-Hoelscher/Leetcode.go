package leetcodeStruct

import (
	"encoding/json"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) ToJSON() string {

	var out, buffer strings.Builder
	out.WriteByte('[')

	var curr []*TreeNode
	if t != nil {
		curr = append(curr, t)
	}

	for len(curr) > 0 {

		var next []*TreeNode

		for _, node := range curr {

			if node != t {
				buffer.WriteByte(',')
			}

			if node == nil {
				buffer.WriteString("null")
				continue
			}

			out.WriteString(buffer.String())
			buffer.Reset()

			out.WriteString(strconv.Itoa(node.Val))
			next = append(next, node.Left, node.Right)

		}

		curr = next
	}

	out.WriteByte(']')
	return out.String()
}

func (t *TreeNode) Equals(o *TreeNode) bool {

	if t == nil || o == nil {
		return t == o
	}

	if t.Val != o.Val {
		return false
	}

	return t.Left.Equals(o.Left) && t.Right.Equals(o.Right)

}

func TreeFromJSON(jsonBlob string) *TreeNode {

	var nums []*int

	if json.Unmarshal([]byte(jsonBlob), &nums) != nil {
		panic("Bad JSON: \"" + jsonBlob + "\"")
	}

	if len(nums) == 0 {
		return nil
	}

	if nums[0] == nil {
		panic("Invalid Tree: \"" + jsonBlob + "\"")
	}

	var nodes []*TreeNode
	for i, x := range nums {

		var newNode *TreeNode
		if x != nil {
			newNode = &TreeNode{*x, nil, nil}
		}

		nodes = append(nodes, newNode)
		parent := nodes[(i-1)/2]

		if i == 0 || parent == nil {
			continue
		}

		if i&1 == 1 {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}
	}

	return nodes[0]
}
