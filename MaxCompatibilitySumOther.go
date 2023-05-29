package main

func maxCompatibilitySumOther(students [][]int, mentors [][]int) int {
	m, n := len(students), len(students[0])
	studentScores := make([]int, m)
	for i := 0; i < m; i++ {
		var scores int
		for j := 0; j < n; j++ {
			scores |= (students[i][j] << (n - j - 1))
		}
		studentScores[i] = scores
	}
	mentorScores := make([]int, m)
	for i := 0; i < m; i++ {
		var scores int
		for j := 0; j < n; j++ {
			scores |= (mentors[i][j] << (n - j - 1))
		}
		mentorScores[i] = scores
	}
	// We have the scores of the students available as bitmasks and we just need
	// to try every possible pairing
	var res int
	mentorTaken := make([]bool, m)
	backtrack(studentScores, mentorScores, n, 0, 0, mentorTaken, &res)
	return res
}

func backtrack(students, mentors []int, q, pos, curr int, mentorTaken []bool, res *int) {
	if pos == len(students) { // All students are matched
		*res = max(*res, curr)
		return
	}
	for j := 0; j < len(mentors); j++ {
		if !mentorTaken[j] { // mentor is available
			mentorTaken[j] = true
			// Calculate compatibility score for current match
			score := calculateScore(students[pos], mentors[j], q)
			backtrack(students, mentors, q, pos+1, curr+score, mentorTaken, res)
			mentorTaken[j] = false
		}
	}
}

func calculateScore(num1, num2 int, bits int) int {
	var res int
	mask := 1
	for j := 0; j < bits; j++ {
		if num1&mask == num2&mask {
			res++
		}
		mask <<= 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
