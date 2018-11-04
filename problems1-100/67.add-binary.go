/* 67. 二进制求和
题目描述提示帮助提交记录社区讨论阅读解答
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101" */

func addBinary(a string, b string) string {
	var carry byte = '0'
	i, j := len(a)-1, len(b)-1
	r := []byte{}
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (a[i] == '0' && (b[j] != carry)) || (a[i] == '1' && (b[j] == carry)) {
			r = append([]byte{'1'}, r...)
		} else {
			r = append([]byte{'0'}, r...)
		}
		if (a[i] == '0' && b[j] == '1' && carry == '1') || (a[i] == '1' && !(b[j] == '0' && carry == '0')) {
			carry = '1'
		} else {
			carry = '0'
		}
	}
	for ; i >= 0; i-- {
		if a[i] == carry {
			r = append([]byte{'0'}, r...)
		} else {
			r = append([]byte{'1'}, r...)
		}
		if a[i] == '1' && carry == '1' {
			carry = '1'
		} else {
			carry = '0'
		}
	}
	for ; j >= 0; j-- {
		if b[j] == carry {
			r = append([]byte{'0'}, r...)
		} else {
			r = append([]byte{'1'}, r...)
		}
		if b[j] == '1' && carry == '1' {
			carry = '1'
		} else {
			carry = '0'
		}
	}
	if carry == '1' {
		r = append([]byte{'1'}, r...)
	}
	return string(r)
}