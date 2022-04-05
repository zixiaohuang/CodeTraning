package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	text1 = " " + text1
	text2 = " " + text2
	var dp [][]int
	for i := 0; i <= m; i++ {
		arr := make([]int, n + 1)
		dp = append(dp, arr)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[m][n]
}

func max(i , j int) int {
	if i < j {
		return j
	}
	return i
}