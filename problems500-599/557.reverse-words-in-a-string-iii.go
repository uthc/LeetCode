/* 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。 */

func reverseWords(s string) string {
	j := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[:j+1] + reverseString(s[j+1:i]) + s[i:]
			j = i
		}
	}
	return s[:j+1] + reverseString(s[j+1:])
}

func reverseString(s string) string {
	str := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}