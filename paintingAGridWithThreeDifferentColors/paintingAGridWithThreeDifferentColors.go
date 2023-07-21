package paintingAGridWithThreeDifferentColors

const mod = 1e9 + 7

func colorTheGrid(m int, n int) int {

	// Find all possible columns of size m.
	var cols [][]int
	hold := make([]int, m)
	var findCols func(i int)
	findCols = func(i int) {
		if i == m {
			newCol := make([]int, m)
			copy(newCol, hold)
			cols = append(cols, newCol)
			return
		}
		for c := 0; c < 3; c++ {
			if i < 1 || hold[i-1] != c {
				hold[i] = c
				findCols(i + 1)
			}
		}
	}
	findCols(0)
	c := len(cols)

	// Helper function to check if col1 allows col2 to be placed after it.
	checkColMatch := func(col1, col2 []int) bool {
		for i := 0; i < m; i++ {
			if col1[i] == col2[i] {
				return false
			}
		}
		return true
	}

	// Make a directed graph from each col1 to valid col2.
	graph := make([][]int, c)
	for i, col1 := range cols {
		for j, col2 := range cols {
			if checkColMatch(col1, col2) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	// Build a memo of the number of valid ways to end with each column.
	memo := make([]int, c)
	for i := range memo {
		memo[i] = 1
	}

	// Iterate across the rest of the width of grid.
	for k := 1; k < n; k++ {
		next := make([]int, c)
		for i, valid := range graph {
			for _, j := range valid {
				next[j] += memo[i]
				next[j] %= mod
			}
		}
		memo = next
	}

	// Sum up the total ways.
	var ans int
	for _, x := range memo {
		ans += x
		ans %= mod
	}

	return ans
}
