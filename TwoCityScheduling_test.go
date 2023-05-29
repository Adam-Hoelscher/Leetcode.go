package main

import (
	"testing"
)

type TwoCitySchedCostInput struct {
	costs [][]int
}

type TwoCitySchedCostCase struct {
	name    string
	input   TwoCitySchedCostInput
	correct int
}

var twoCitySchedCostTestCases []TwoCitySchedCostCase = []TwoCitySchedCostCase{
	{
		name:    "Example 1",
		input:   TwoCitySchedCostInput{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}},
		correct: 110,
	},
	{
		name:    "Example 2",
		input:   TwoCitySchedCostInput{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}},
		correct: 1859,
	},
	{
		name:    "Example 3",
		input:   TwoCitySchedCostInput{[][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}},
		correct: 3086,
	},
	{
		name:    "Example 49",
		input:   TwoCitySchedCostInput{[][]int{{115, 14}, {932, 665}, {528, 396}, {633, 293}, {948, 428}, {757, 520}, {932, 620}, {590, 142}, {157, 353}, {218, 858}, {394, 121}, {927, 729}, {737, 722}, {94, 845}, {567, 360}, {83, 123}, {320, 323}, {935, 893}, {303, 142}, {6, 146}, {668, 346}, {234, 625}, {617, 616}, {550, 105}, {949, 691}, {923, 876}, {569, 300}, {558, 768}, {381, 732}, {288, 541}, {764, 829}, {37, 913}, {395, 25}, {606, 959}, {974, 58}, {31, 635}, {516, 880}, {113, 480}, {310, 959}, {185, 618}, {516, 470}, {951, 940}, {744, 328}, {808, 676}, {808, 487}, {53, 186}, {728, 700}, {202, 973}, {614, 200}, {552, 709}, {589, 183}, {43, 474}, {82, 960}, {388, 804}, {930, 118}, {800, 392}, {770, 306}, {191, 2}, {333, 630}, {462, 917}, {146, 459}, {197, 469}, {324, 500}, {90, 965}, {8, 169}, {381, 217}, {26, 557}, {602, 286}, {136, 321}, {298, 669}, {357, 127}, {897, 596}, {16, 335}, {561, 828}, {162, 213}, {982, 652}, {428, 497}, {265, 548}, {703, 805}, {821, 755}, {937, 643}, {901, 674}, {47, 621}, {340, 71}, {557, 139}, {147, 539}, {896, 304}, {526, 348}, {413, 909}, {239, 28}, {743, 269}, {930, 672}, {13, 183}, {256, 393}, {705, 464}, {20, 71}, {282, 182}, {545, 137}, {411, 86}, {105, 78}}},
		correct: 31918,
	},
}

func TestTwoCitySchedCost(t *testing.T) {

	for _, testCase := range twoCitySchedCostTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := twoCitySchedCost(testCase.input.costs)
			if actual != testCase.correct {
				t.Error(actual, "should be", testCase.correct)
			}
		})
	}

}
