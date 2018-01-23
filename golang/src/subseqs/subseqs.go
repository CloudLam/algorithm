// Copyright 2018 Cloud Lam. All rights reserved
// Package subseqs implements some subsequence methods

package main

import (
	"fmt"
)

// LCS returns Longest Common Subsequence
func LCS(x, y string) (int, string) {
	if len(x) == 0 || len(y) == 0 {
		return 0, ""
	}
	lcs := ""
	var dp [][]int
	for i := 0; i <= len(x); i++ {
		var temp []int
		for j := 0; j <= len(y); j++ {
			if i == 0 {
				temp = append(temp, 0)
				continue
			}
			if x[i-1] == y[j-1] {
				temp = append(temp, dp[i-1][j-1]+1)
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					temp = append(temp, dp[i-1][j])
				} else {
					temp = append(temp, dp[i][j-1])
				}
			}
		}
		dp = append(dp, temp)
		fmt.Println(dp)
	}
	return dp[len(x)][len(y)], lcs
}

func main() {
	len, lcs := LCS("lcs", "lcs1")
	fmt.Println(len, lcs)
}

/*var dp = [new Array(y.length + 1).fill(0)];
  var lcs = '';
  // Length of LCS
  for (var i = 1; i <= x.length + 1; i++) {
    dp[i] = [0];
    for (var j = 1; j <= y.length; j++) {
      if (x[i-1] == y[j-1]) {
        dp[i][j] = dp[i-1][j-1] + 1;
      } else {
        dp[i][j] = Math.max(dp[i-1][j], dp[i][j-1]);
      }
    }
  }
  // Subsequence of LCS
  for (var m = x.length, n = y.length; m > 0 && n > 0;) {
    if (x[m-1] == y[n-1]) {
      lcs = x[m-1] + lcs;
      m--;
      n--;
      continue
    }
    if (dp[m][n-1] > dp[m-1][n]) {
      n--;
    } else {
      m--;
    }
  }*/
