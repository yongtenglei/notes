# 最长回文序列 [longestPalindrome](https://leetcode-cn.com/problems/longest-palindromic-substring/)

## 动态规划解法

由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串. 所以, 当确定长字串是否为回文字串时, 可以向内推导, 即 `dp[i][j] = dp[i+1][j-1])`, 如果里面的子串为回文字串, 则此字串为回文子串.

状态转移方程:

<div align=center>dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1]</div>

### 从后往前推

一个 flag i 指向最后一个字符, 另一个 flag j 从 i 开始向后判断 i 到 j 的字串是否为回文子串.

- 如果 i 与 j 字符不相等

  - 则不为回文子串. 即`dp[i][j] = false`

- 如果 i 与 j 字符相等
  - 如果 j - i < 3, 则代表中间只有一个字符, 为回文子串. 即`dp[i][j] = true`
  - 向内推导, 判断是否为回文子串. 即 `dp[i][j] = dp[i+1][j-1])`

```go
func longestPalindrome(s string) string {
	res := ""

	// 初始化 dp[i][j] 代表i 到 j 的字串是否为回文子串
	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s), len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 如果i 到 j 的子串为回文子串, 且第一次更新或者比现有的res长
			// j - i + 1 为子串的长度
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}

	return res
}

```

### 从前往后推

一个 flag j 指向第一个字符, 另一个 flag i 从第一个字符开始判断 i 到 j 之前的字串是否为回文子串.

```go
func longestPalindrome2(s string) string {
	res := ""

	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s), len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 如果i 到 j 的子串为回文子串, 且第一次更新或者比现有的res长
			// j - i + 1 为子串的长度
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}

	}

	return res

}

```
