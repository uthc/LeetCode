/* 检测大写字母
给定一个单词，你需要判断单词的大写使用是否正确。

我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如"USA"。
单词中所有字母都不是大写，比如"leetcode"。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
否则，我们定义这个单词没有正确使用大写字母。

示例 1:

输入: "USA"
输出: True
示例 2:

输入: "FlaG"
输出: False
注意: 输入是由大写和小写拉丁字母组成的非空单词。 */

func detectCapitalUse(word string) bool {
	l := len(word)
	if l <= 1 {
		return true
	}
	if word[0] >= 'a' {
		i := 1
		for ; i < l && word[i] >= 'a'; i++ {
		}
		return i == l
	}
	if word[l-1] >= 'a' {
		i := l - 1
		for ; i >= 0 && word[i] >= 'a'; i-- {
		}
		return i == 0
	}
	if word[l-1] <= 'Z' {
		i := l - 1
		for ; i >= 0 && word[i] <= 'Z'; i-- {
		}
		return i == -1
	}
	return false
}