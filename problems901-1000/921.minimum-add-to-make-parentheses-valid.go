import "fmt"

/* 921. 使括号有效的最少添加
用户通过次数 172
用户尝试次数 184
通过次数 173
提交次数 269
题目难度 Medium
给定一个由 '(' 和 ')' 括号组成的字符串 S，我们需要添加最少的括号（ '(' 或是 ')'，可以在任何位置），以使得到的括号字符串有效。

从形式上讲，只有满足下面几点之一，括号字符串才是有效的：

它是一个空字符串，或者
它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
它可以被写作 (A)，其中 A 是有效字符串。
给定一个括号字符串，返回为使结果字符串有效而必须添加的最少括号数。



示例 1：

输入："())"
输出：1
示例 2：

输入："((("
输出：3
示例 3：

输入："()"
输出：0
示例 4：

输入："()))(("
输出：4


提示：

S.length <= 1000
S 只包含 '(' 和 ')' 字符。 */
func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}
	stack := []rune{}
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) == 0 {
				stack = append(stack, ')')
			} else if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, ')')
				fmt.Println(stack)
			}
		}
	}
	return len(stack)
}