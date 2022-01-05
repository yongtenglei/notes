package main

import "fmt"

// longestPalindrome 从后往前推
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

// longestPalindrome2 从前往后推
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
func main() {
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome2("aacabdkacaa"))
}
