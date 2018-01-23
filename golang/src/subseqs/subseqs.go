// Copyright 2018 Cloud Lam. All rights reserved
// Package subseqs implements some subsequence methods

package subseqs

// LCS returns Longest Common Subsequence
func LCS(x, y string) (int, string) {
	if len(x) == 0 || len(y) == 0 {
		return 0, ""
	}
	lcs := ""
	var dp [][]int
	// Length of LCS
	for i := 0; i <= len(x); i++ {
		var temp []int
		for j := 0; j <= len(y); j++ {
			if i == 0 || j == 0 {
				temp = append(temp, 0)
				continue
			}
			if x[i-1] == y[j-1] {
				temp = append(temp, dp[i-1][j-1]+1)
			} else {
				if dp[i-1][j] > temp[j-1] {
					temp = append(temp, dp[i-1][j])
				} else {
					temp = append(temp, temp[j-1])
				}
			}
		}
		dp = append(dp, temp)
	}
	// Subsequence of LCS
	m := len(x)
	n := len(y)
	for m > 0 && n > 0 {
		if x[m-1] == y[n-1] {
			lcs = string(x[m-1]) + lcs
			m--
			n--
			continue
		}
		if dp[m][n-1] > dp[m-1][n] {
			n--
		} else {
			m--
		}
	}
	return dp[len(x)][len(y)], lcs
}
