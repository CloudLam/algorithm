// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"subseqs"
)

// LCSTest tests Longest Common Subsequence
func LCSTest() {
	len, lcs := subseqs.LCS("lcs", "longest-common-subsequence")
	fmt.Println(len, lcs)
}
