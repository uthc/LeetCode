/* 66. 加一
题目描述提示帮助提交记录社区讨论阅读解答
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。 */

func plusOne(digits []int) []int {
	l := len(digits)
	add := false
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			add = true
			break
		} else {
			digits[i] = 0
		}
	}
	if !add {
		s := []int{1}
		digits = append(s, digits...)
	}
	return digits
}