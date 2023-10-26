package paintingAGridWithThreeDifferentColors

import (
	// . "leetcode/leetcodeStruct"

	"fmt"
	"sort"
	"testing"
)

type ColorTheGridCase struct {
	name    string
	m       int
	n       int
	correct int
}

var colorTheGridTestCases []ColorTheGridCase = []ColorTheGridCase{
	{name: "1*1", m: 1, n: 1, correct: 3},
	{name: "1*2", m: 1, n: 2, correct: 6},
	{name: "5*5", m: 5, n: 5, correct: 580_986},
	{name: "5*1000", m: 5, n: 1000, correct: 408_208_448},
}

var colorTheGridBenchCases []ColorTheGridCase

func init() {
	for n := 256; n <= 1024; n *= 2 {
		for m := 15; m <= 15; m++ {
			testCase := ColorTheGridCase{
				name:    fmt.Sprintf("%v*%v", m, n),
				m:       m,
				n:       n,
				correct: 0,
			}
			colorTheGridBenchCases = append(colorTheGridBenchCases, testCase)
		}
	}
	sort.Slice(colorTheGridBenchCases, func(i, j int) bool {
		if colorTheGridBenchCases[i].n != colorTheGridBenchCases[j].n {
			return colorTheGridBenchCases[i].n < colorTheGridBenchCases[j].n
		}
		return colorTheGridBenchCases[i].m < colorTheGridBenchCases[j].m
	})
}

func TestColorTheGrid(t *testing.T) {

	for _, testCase := range colorTheGridTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := colorTheGrid(testCase.m, testCase.n)
			if actual != testCase.correct {
				t.Error(actual, "should be", testCase.correct)
			}
		})
	}

}

func BenchmarkColorTheGrid(b *testing.B) {

	for _, testCase := range colorTheGridBenchCases {
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				colorTheGrid(testCase.m, testCase.n)
			}
			b.ReportMetric(b.Elapsed().Seconds()/float64(b.N), "s/op")
		})
	}

}
