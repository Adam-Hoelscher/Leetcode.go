package main

import (
	"testing"
)

type MinDifficultyJobSchedInput struct {
	jobDifficulty []int
	d             int
}

type MinDifficultyJobSchedCase struct {
	name    string
	input   MinDifficultyJobSchedInput
	correct int
}

var minDifficultyTestCases []MinDifficultyJobSchedCase = []MinDifficultyJobSchedCase{
	{
		name: "Example 1",
		input: MinDifficultyJobSchedInput{
			[]int{6, 5, 4, 3, 2, 1},
			2,
		},
		correct: 7,
	},
	{
		name: "Example 2",
		input: MinDifficultyJobSchedInput{
			[]int{9, 9, 9},
			4,
		},
		correct: -1,
	},
	{
		name: "Example 3",
		input: MinDifficultyJobSchedInput{
			[]int{1, 1, 1},
			3,
		},
		correct: 3,
	},
	{
		name: "Example 34",
		input: MinDifficultyJobSchedInput{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			10,
		},
		correct: 0,
	},
}

func TestMinDifficulty(t *testing.T) {

	for _, testCase := range minDifficultyTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := minDifficulty(testCase.input.jobDifficulty, testCase.input.d)
			if actual != testCase.correct {
				t.Error(actual, "should be", testCase.correct)
			}
		})
		break
	}

}
