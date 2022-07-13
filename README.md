# golang_edit_distance

Given two strings `word1` and `word2`, return *the minimum number of operations required to convert `word1` to `word2`*.

You have the following three operations permitted on a word:

- Insert a character
- Delete a character
- Replace a character

## Examples

**Example 1:**

```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

```

**Example 2:**

```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

```

**Constraints:**

- `0 <= word1.length, word2.length <= 500`
- `word1` and `word2` consist of lowercase English letters.

## 解析

給兩個字串 word1, word2

如果要把 word1 轉換成 word2

可以做以下三種運算

1 取代 word1 的一個字元

2 刪除 word1 的一個字元

3 新增一個字元到 word1

要求寫一個演算法算出 從 word1 轉換到 word2 的最小步驟

如果直接對每次修改的 word1 做操作 再用操做過 word1 來比對會讓整個思考變得太複雜

所以每次比較 都原本的 word1 word2 相對位置字元做思考

![](https://i.imgur.com/sXLyww2.png)

透過以 (i,j) 表示從 word1 i 位置, word2 j 開始比較 可以畫出以下決策樹

 

![](https://i.imgur.com/ZUGiwx8.png)

可以發現

對每個 i,j 都有三種選項

要到 len(word1), len(word2) 才會結束比較

所以要走訪所有可能必須透過 DFS 走完所有結點

這樣會是 $3^{m*n}$, where m =len(word1), n=len(word2)

透過 cache 可以把走過的結點都暫存下來避免重複走訪

因為每個起始點會有 m by n

所以需要 時間複雜度 是 O(m*n)

而空間複雜度是 O(m*n)

如果是透過遞迴 DFS 則 call stack 的空間複雜度也是 O(m*n)

因為要到 (m,n) 才會被 resolve

透過  Tabulation Dynamic Programming 則可以減少 call stack 的消耗

一樣的透過前面比較的方式

可以定義 dp[i][j] = 代表 word1[i:] 轉換成 word2[j:] 的最少步驟

可發現有以下關係

 
![](https://i.imgur.com/ccKSrEi.png)

以下有幾個 edge case 要思考一下

假設 是當 word1 是 空字串時, 則要修改的最少步驟會是 把 word2 的所有字元加入

所以是 len(word2)

同樣的當 word2 是空字串時，則要修改的最少步驟會是 把 word1 的所有字元刪除

所以是 len(word1)

而當兩個字串都是 空字串時 則不需要做任何步驟

所求會是 dp[0][0] 從 word1[0:] 轉換成 word2[0:]

時間複雜度是 O(m*n) where m=len(word1), n= len(word2)

空間複雜度是 O(m*n) where m=len(word1), n= len(word2)

 但比遞迴DFS 減少了 stack space

## 程式碼
```go
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

```
## 困難點

1. 要看出兩個字串起始位置與最少步驟的相對關係

## Solve Point

- [x]  初始化一個整數矩陣 dp 大小為 len(word1) by len(word2)
- [x]  初始化 dp[len(word1)][j] = len(word2) - j , dp[i][len(word2)] = len(word1) - i
- [x]  從 i = len(word1) -1 , j = len(word2) - 1 開始逐步透過關係式計算每個 dp[i][j]
- [x]  dp[0][0] 及為所求