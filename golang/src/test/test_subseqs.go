// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"regexp"
	"subseqs"
)

// LCSTest tests Longest Common Subsequence
func LCSTest() {
	x := "lcs"
	y := "longest-common-subsequence"

	length, lcs := subseqs.LCS(x, y)

	fmt.Println(length, lcs)

	re := string(lcs[0])
	for i := 1; i < len(lcs); i++ {
		re += ".*" + string(lcs[i])
	}
	matchX, _ := regexp.Match(re, []byte(x))
	matchY, _ := regexp.Match(re, []byte(y))

	fmt.Println(re, matchX && matchY)
}
