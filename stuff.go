package main

func shortestPath(grid [][]int, k int) int {

	m := len(grid)
	n := len(grid[0])

	maxObstacles := make([][]int, m)
	for i := range maxObstacles {
		maxObstacles[i] = make([]int, n)
		for j := range maxObstacles[i] {
			maxObstacles[i][j] = -1
		}
	}

	queue := [][]int{{0, 0, k}}
	maxObstacles[0][0] = k

	var next [][]int
	for steps := 0; len(queue) > 0; steps++ {

		for _, state := range queue {

			i, j, r := state[0], state[1], state[2]

			r -= grid[i][j]
			if r < 0 {
				continue
			}

			if i == m-1 && j == n-1 {
				return steps
			}

			if i > 0 && r > maxObstacles[i-1][j] {
				next = append(next, []int{i - 1, j, r})
				maxObstacles[i-1][j] = r
			}

			if i+1 < m && r > maxObstacles[i+1][j] {
				next = append(next, []int{i + 1, j, r})
				maxObstacles[i+1][j] = r
			}

			if j > 0 && r > maxObstacles[i][j-1] {
				next = append(next, []int{i, j - 1, r})
				maxObstacles[i][j-1] = r
			}

			if j+1 < n && r > maxObstacles[i][j+1] {
				next = append(next, []int{i, j + 1, r})
				maxObstacles[i][j+1] = r
			}
		}

		queue, next = next, queue[:0]
	}

	return -1
}
