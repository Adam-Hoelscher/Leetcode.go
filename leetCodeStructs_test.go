package main

import (
	"testing"
)

func TestTreeToJSON(t *testing.T) {

	var tree *TreeNode
	var actual, correct string

	t.Run("Empty", func(t *testing.T) {
		tree = nil
		correct = "[]"
		actual = tree.ToJSON()
		if actual != correct {
			t.Error(actual, "should be", correct)
		}
	})

	t.Run("Left", func(t *testing.T) {
		tree = &TreeNode{0, &TreeNode{Val: 1}, nil}
		correct = "[0,1]"
		actual = tree.ToJSON()
		if actual != correct {
			t.Error(actual, "should be", correct)
		}
	})

	t.Run("Right", func(t *testing.T) {
		tree = &TreeNode{0, nil, &TreeNode{Val: 2}}
		correct = "[0,null,2]"
		actual = tree.ToJSON()
		if actual != correct {
			t.Error(actual, "should be", correct)
		}
	})

	t.Run("Both", func(t *testing.T) {
		tree = &TreeNode{0, &TreeNode{Val: 1}, &TreeNode{Val: 2}}
		correct = "[0,1,2]"
		actual = tree.ToJSON()
		if actual != correct {
			t.Error(actual, "should be", correct)
		}
	})

	t.Run("Large", func(t *testing.T) {
		tree = &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 6},
			},
		}
		correct = "[0,1,2,3,4,null,6,7]"
		actual = tree.ToJSON()
		if actual != correct {
			t.Error(actual, "should be", correct)
		}
	})

}

func TestTreeFromJSON(t *testing.T) {

	var jsonBlob string
	var actual, correct *TreeNode

	t.Run("Empty", func(t *testing.T) {
		jsonBlob = "[]"
		correct = nil
		actual = TreeFromJSON(jsonBlob)
		if !actual.Equals(correct) {
			t.Error(actual.ToJSON(), "should be", correct.ToJSON())
		}
	})

	t.Run("Left", func(t *testing.T) {
		jsonBlob = "[0,1]"
		correct = &TreeNode{0, &TreeNode{Val: 1}, nil}
		actual = TreeFromJSON(jsonBlob)
		if !actual.Equals(correct) {
			t.Error(actual.ToJSON(), "should be", correct.ToJSON())
		}
	})

	t.Run("Right", func(t *testing.T) {
		jsonBlob = "[0,null,2]"
		correct = &TreeNode{0, nil, &TreeNode{Val: 2}}
		actual = TreeFromJSON(jsonBlob)
		if !actual.Equals(correct) {
			t.Error(actual.ToJSON(), "should be", correct.ToJSON())
		}
	})

	t.Run("Both", func(t *testing.T) {
		jsonBlob = "[0,1,2]"
		correct = &TreeNode{0, &TreeNode{Val: 1}, &TreeNode{Val: 2}}
		actual = TreeFromJSON(jsonBlob)
		if !actual.Equals(correct) {
			t.Error(actual.ToJSON(), "should be", correct.ToJSON())
		}
	})

	t.Run("Large", func(t *testing.T) {
		jsonBlob = "[0,1,2,3,4,null,6,7]"
		correct = &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 6},
			},
		}
		actual = TreeFromJSON(jsonBlob)
		if !actual.Equals(correct) {
			t.Error(actual.ToJSON(), "should be", correct.ToJSON())
		}
	})

}
