package paintingAGridWithThreeDifferentColors

const mod = 1e9 + 7
const colors = 3

func colorTheGrid(m int, n int) int {

	// Find all possible columns of size m.
	cols := [][]int{{0}}

	for k, next := 1, [][]int(nil); k < m; k++ {
		for _, poss := range cols {
			for s := 1; s < colors; s++ {
				add := make([]int, len(poss)+1)
				copy(add, poss)
				add[k] = (add[k-1] + s) % colors
				next = append(next, add)
			}
		}
		cols, next = next, cols[:0]
	}
	c := len(cols)

	// Helper function to count rotations of col2 that can come after col1.
	countColMatch := func(col1, col2 []int) int {

		var i, weight int

		for s := 1; s < colors; s++ {
			for i = 0; i < m; i++ {
				if col1[i] == (col2[i]+s)%colors {
					break
				}
			}
			if i == m {
				weight++
			}
		}

		return weight
	}

	// Make a directed graph from each col1 to valid col2.
	graph := make([][]int, c)
	for i, col1 := range cols {
		graph[i] = make([]int, c)
		for j := range cols[:i] {
			graph[i][j] = graph[j][i]
		}
		for j, col2 := range cols[i:] {
			graph[i][i+j] = countColMatch(col1, col2)
		}
	}

	// Build a memo of the number of valid ways to end coloring for each column.
	memo := make([]int, c)
	for i := range memo {
		memo[i] = colors
	}

	// Iterate across the rest of the width of grid.
	for k := 1; k < n; k++ {
		next := make([]int, c)
		for i, valid := range graph {
			for j, weight := range valid {
				next[j] += weight * memo[i]
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
