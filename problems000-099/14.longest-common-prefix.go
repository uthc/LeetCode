/*  最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。 */

func longestCommonPrefix(strs []string) string {
	s := ""
	x, count := len(strs), 0
	if x == 0 {
		return s
	} else if x == 1 {
		return strs[0]
	}
	for j := true; j; count++ {
		for i := 0; i < x; i++ {
			if count < len(strs[i]) && strs[i][count] != strs[0][count] {
				j = false
				break
			}
			if count == len(strs[i]) {
				j = false
				break
			}
		}
	}
	s = string(strs[0][0 : count-1])
	return s
}