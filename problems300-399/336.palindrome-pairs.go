/* 336.回文对
给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1:

输入: ["abcd","dcba","lls","s","sssll"]
输出: [[0,1],[1,0],[3,2],[2,4]]
解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2:

输入: ["bat","tab","cat"]
输出: [[0,1],[1,0]]
解释: 可拼接成的回文串为 ["battab","tabbat"] */

func palindromePairs(words []string) [][]int {
	result := [][]int{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if isPalindrome(words[i], words[j]) {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func isPalindrome(s1, s2 string) bool {
	x := []byte(s1)
	x = append(x, []byte(s2)...)
	for h, t := 0, len(x)-1; h < t; h, t = h+1, t-1 {
		if x[h] != x[t] {
			return false
		}
	}
	return true
}