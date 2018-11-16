/* 反转字符串
编写一个函数，其作用是将输入的字符串反转过来。

示例 1:

输入: "hello"
输出: "olleh"
示例 2:

输入: "A man, a plan, a canal: Panama"
输出: "amanaP :lanac a ,nalp a ,nam A" */

func reverseString(s string) string {
	// if len(s) < 2 {
	// 	return s
	// }
	// return string(s[len(s)-1]) + reverseString(string(s[1:len(s)-1])) + string(s[0])
	str := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}