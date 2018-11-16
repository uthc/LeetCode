/* 字符串中的单词数
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

示例:

输入: "Hello, my name is John"
输出: 5 */

func countSegments(s string) int {
	count := 0
	for i, v := range s {
		if v == ' ' && i > 0 && s[i-1] != ' ' {
			count++
		}
	}
	if len(s) > 0 && s[len(s)-1] != ' ' {
		count++
	}
	return count
}