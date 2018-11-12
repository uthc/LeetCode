/* 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb" */

func longestPalindrome(s string) string {
	start := 0
	end := 0
	count := 0
	length := len(s)
	if length > 0 {
		end = 1
		count = 1
	}
	for i := 0; i < length; i++ {

		if i < length-1 && s[i] == s[i+1] {
			m := i
			n := i + 1
			for m > 0 && n < length-1 && s[m-1] == s[n+1] {
				m--
				n++
			}
			if count < n-m+1 {
				count = n - m + 1
				start = m
				end = n + 1
			}
		}
		if i > 0 && i < length-1 && s[i-1] == s[i+1] {
			m := i
			n := i
			for m > 0 && n < length-1 && s[m-1] == s[n+1] {
				m--
				n++
			}
			if count < n-m+1 {
				count = n - m + 1
				start = m
				end = n + 1
			}

		}
	}
	result := s[start:end]
	return result
}