package tests

import (
	"fmt"
)

// Adapted from Python implementation by Florian Hartmann, https://github.com/florian/diff-tool
func Diff(a []string, b []string) (results []string, textsEqual bool) {
	textsEqual = true
	lcs := lcsLen(a, b)

	i := len(a)
	j := len(b)

	for i != 0 || j != 0 {
		if i == 0 {
			textsEqual = false
			results = append(results, fmt.Sprintf("+ %s", b[j - 1]))
			j -= 1
		} else if j == 0 {
			textsEqual = false
			results = append(results, fmt.Sprintf("- %s", a[i - 1]))
			i -= 1
		} else if a[i - 1] == b[j - 1] {
			results = append(results, fmt.Sprintf("  %s", a[i - 1]))
			i -= 1
			j -= 1
		} else if lcs[i - 1][j] <= lcs[i][j - 1] {
			textsEqual = false
			results = append(results, fmt.Sprintf("+ %s", b[j - 1]))
			j -= 1
		} else {
			textsEqual = false
			results = append(results, fmt.Sprintf("- %s", a[i - 1]))
			i -= 1
		}
	}

	if textsEqual {
		return
	}

	// Reverse results
	for i := 0; i < len(results) / 2; i++ {
		temp := results[i]
		results[i] = results[len(results) - 1 - i]
		results[len(results) - 1 - i] = temp
	}
	return
}

func lcsLen(a []string, b []string) [][]int {
	n := len(a)
	m := len(b)

	lcs := make([][]int, n + 1)
	for r := 0; r < len(lcs); r++ {
		lcs[r] = make([]int, m + 1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if a[i - 1] == b[j - 1] {
				lcs[i][j] = 1 + lcs[i - 1][j - 1]
			} else {
				lcs[i][j] = max(lcs[i - 1][j], lcs[i][j - 1])
			}
		}
	}

	return lcs
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
