/* 回文数
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？ */

func isPalindrome(x int) bool {
	var judge bool = true
	c := 0
	j := 0
	if x < 0 {
		judge = false
	} else {
		count := 0
		n := 1
		for i := x; i != 0; i = i / 10 {
			count++
			n = n * 10
		}
		n = n / 10
		for x/10 != 0 {
			if x%10 != x/n {
				judge = false
				break
			}
			x = x % n
			x = x / 10
			n = n / 100
			count = count - 2
			c = count
			j++
			if x == 0 {
				judge = true
				j = 0
				break
			}
		}
		if c != 0 && c != 1 && j != 0 {
			judge = false
		}
	}
	return judge
}