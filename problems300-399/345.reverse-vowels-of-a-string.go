/* 反转字符串中的元音字母
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede" */

func reverseVowels(s string) string {
	f := func(b byte) bool {
		if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
			return false
		}
		return true
	}
	str := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		if f(s[i]) {
			i++
		}
		if f(s[j]) {
			j--
		}
		if !f(s[i]) && !f(s[j]) {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}