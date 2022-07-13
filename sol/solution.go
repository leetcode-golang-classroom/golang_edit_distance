package sol

func minDistance(word1 string, word2 string) int {
	word1Len := len(word1)
	word2Len := len(word2)
	// dp[i][j]: min step number from word1[i:] to word2[j:]
	dp := make([][]int, word1Len+1)
	for row := range dp {
		dp[row] = make([]int, word2Len+1)
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// initial edge case, word1 = "" or word2 = ""
	for word1Start := 0; word1Start <= word1Len; word1Start++ {
		dp[word1Start][word2Len] = word1Len - word1Start
	}
	for word2Start := 0; word2Start <= word2Len; word2Start++ {
		dp[word1Len][word2Start] = word2Len - word2Start
	}
	for word1Start := word1Len - 1; word1Start >= 0; word1Start-- {
		for word2Start := word2Len - 1; word2Start >= 0; word2Start-- {
			if word1[word1Start] == word2[word2Start] {
				dp[word1Start][word2Start] = dp[word1Start+1][word2Start+1]
			} else {
				dp[word1Start][word2Start] = 1 + min(
					dp[word1Start+1][word2Start+1],
					min(dp[word1Start+1][word2Start], dp[word1Start][word2Start+1]),
				)
			}
		}
	}
	return dp[0][0]
}
