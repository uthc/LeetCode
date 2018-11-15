/* 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true */

func isValid(s string) bool {
	stack := make([]string, 0)
	p := 0
	var ops map[string]string
	ops = make(map[string]string)
	ops[")"] = "("
	ops["]"] = "["
	ops["}"] = "{"
	for i := 0; i < len(s); i++ {
		t := string(s[i])
		if t == "(" || t == "[" || t == "{" {
			stack = append(stack, t)
			p++
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[p-1] == ops[t] {
				stack = stack[:p-1]
				p--
			} else {
				return false
			}
		}
	}
	if p == 0 {
		return true
	} else {
		return false
	}
}