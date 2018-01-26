/*!
 * subseq.js v1.0.0
 * 2017-2018 Cloud Lam
 * Released under the MIT License.
 */

 /**
 * Longest Common Subsequence
 * 
 * @param x
 * @param y
 * @return {len, lcs}
 */
function LCSeq (x, y) {
  if (x.length == 0 || y.length == 0) {
    return {len: 0, lcs: ''};
  }
  var dp = [new Array(y.length + 1).fill(0)];
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
  }
  return {len: dp[x.length][y.length], lcs: lcs};
}

 /**
 * Longest Common Substring
 * 
 * @param x
 * @param y
 * @return {len, lcs}
 */
function LCStr (x, y) {
  if (x.length == 0 || y.length == 0) {
    return {len: 0, lcs: ''};
  }
  var dp = [new Array(y.length + 1).fill(0)];
  var len = 0;
  var lcs = '';
  for (var i = 1; i <= x.length + 1; i++) {
    dp[i] = [0];
    for (var j = 1; j <= y.length; j++) {
      if (x[i-1] == y[j-1]) {
        dp[i][j] = dp[i-1][j-1] + 1;
        if (dp[i][j] > len) {
          // Length of LCS
          len = dp[i][j];
          // Substring of LCS
          lcs = '';
          for (var k = len; k > 0; k--) {
            console.log(y[j-k]);
            lcs += y[j-k];
          }
        }
      } else {
        dp[i][j] = 0;
      }
    }
  }
  return {len: len, lcs: lcs};
}

 /**
 * Longest Increasing Subsequence
 * 
 * @param x
 * @return {len, lis}
 */
function LIS (x) {
  return {len: 0, lis: ''};
}
