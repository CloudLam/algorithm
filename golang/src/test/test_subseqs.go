// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"regexp"
	"subseqs"
)

// SubseqTest tests subseq methods
func SubseqTest() {
	LCSeqTest()
	LCStrTest()
	LISeqTest()
	LIStrTest()
}

// LCSeqTest tests Longest Common Subsequence
func LCSeqTest() {
	x := "lcs"
	y := "longest-common-subsequence"

	length, lcs := subseqs.LCSeq(x, y)

	fmt.Print("LCS(subsequence), Length: ", length, "    LCS: ", lcs)

	re := string(lcs[0])
	for i := 1; i < len(lcs); i++ {
		re += ".*" + string(lcs[i])
	}
	matchX, _ := regexp.Match(re, []byte(x))
	matchY, _ := regexp.Match(re, []byte(y))

	fmt.Println("    RE:", re, "    Test:", matchX && matchY)
}

// LCStrTest tests Longest Common Substring
func LCStrTest() {
	x := "lcs"
	y := "longest-common-substring"

	length, lcs := subseqs.LCStr(x, y)

	fmt.Print("LCS(substring), Length: ", length, "    LCS: ", lcs)

	re := lcs + ".*"
	matchX, _ := regexp.Match(re, []byte(x))
	matchY, _ := regexp.Match(re, []byte(y))

	fmt.Println("    RE:", re, "    Test:", matchX && matchY)
}

// LISeqTest tests Longest Increasing Subsequence
func LISeqTest() {
	x := "31264597"

	length, lis := subseqs.LISeq(x)

	fmt.Println("LIS(subsequence), Length: ", length, "    LIS: ", lis)
}

// LIStrTest tests Longest Increasing Substring
func LIStrTest() {
	x := "31264597"

	length, lis := subseqs.LIStr(x)

	fmt.Println("LIS(substring), Length: ", length, "    LIS: ", lis)
}
