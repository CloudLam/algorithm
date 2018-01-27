// Copyright 2018 Cloud Lam. All rights reserved
// Package subseqs implements some subsequence methods

package subseqs

// LCSeq returns Longest Common Subsequence: length & 1 of LCS
func LCSeq(x, y string) (int, string) {
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

// LCStr returns Longest Common Substring: length & 1 of LCS
func LCStr(x, y string) (int, string) {
	if len(x) == 0 || len(y) == 0 {
		return 0, ""
	}
	length := 0
	lcs := ""
	var dp [][]int
	for i := 0; i <= len(x); i++ {
		var temp []int
		for j := 0; j <= len(y); j++ {
			if i == 0 || j == 0 {
				temp = append(temp, 0)
				continue
			}
			if x[i-1] == y[j-1] {
				temp = append(temp, dp[i-1][j-1]+1)
				if temp[j] > length {
					// Length of LCS
					length = temp[j]
					// Substring of LCS
					lcs = ""
					for k := length; k > 0; k-- {
						lcs += string(y[j-k])
					}
				}
			} else {
				temp = append(temp, 0)
			}
		}
		dp = append(dp, temp)
	}
	return length, lcs
}

// LISeq returns Longest Increasing Subsequence: length & 1 of LIS
func LISeq(x string) (int, string) {
	return 0, ""
}

// LIStr returns Longest Increasing Substring: length & 1 of LIS
func LIStr(x string) (int, string) {
	if len(x) == 0 {
		return 0, ""
	}
	length := 1
	lis := ""
	var dp []int
	// Length of LIS
	dp = append(dp, 1)
	for i := 1; i < len(x); i++ {
		if x[i] > x[i-1] {
			dp = append(dp, dp[i-1]+1)
		} else {
			dp = append(dp, 1)
		}
		if dp[i] > length {
			length = dp[i]
		}
	}
	// Substring of LIS
	for i := 0; i < len(dp); i++ {
		if dp[i] == length {
			for j := length - 1; j >= 0; j-- {
				lis += string(x[i-j])
			}
			break
		}
	}
	return length, lis
}
