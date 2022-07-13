package sol

func minDistanceDFS(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	cache := make([][]int, word1Len)
	for row := range cache {
		cache[row] = make([]int, word2Len)
	}
	for word1End := 0; word1End < word1Len; word1End++ {
		for word2End := 0; word2End < word2Len; word2End++ {
			cache[word1End][word2End] = -1
		}
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(word1End, word2End int) int
	dfs = func(word1End, word2End int) int {
		if word1End == word1Len && word2End == word2Len {
			return 0
		}
		if word1End == word1Len { // word1 use all
			return word2Len - word2End
		}
		if word2End == word2Len { // word2 use all
			return word1Len - word1End
		}
		if cache[word1End][word2End] != -1 {
			return cache[word1End][word2End]
		}
		var result int
		if word1[word1End] == word2[word2End] {
			result = dfs(word1End+1, word2End+1)
		} else {
			result = 1 + min(
				dfs(word1End+1, word2End+1),
				min(dfs(word1End+1, word2End),
					dfs(word1End, word2End+1),
				),
			)
		}
		cache[word1End][word2End] = result
		return result
	}
	return dfs(0, 0)
}
