/* Z字形变换
将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：

P   A   H   N
A P L S I I G
Y   I   R
之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

实现一个将字符串进行指定行数变换的函数:

string convert(string s, int numRows);
示例 1:

输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
示例 2:

输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:

P     I    N
A   L S  I G
Y A   H R
P     I */

func convert(s string, numRows int) string {
	sl := len(s)
	result := make([]byte, sl)
	r := 0
	for i := 0; i < numRows; i++ {
		x1 := 2 * (numRows - i - 1)
		x2 := 2 * (i)
		if numRows == 1 {
			x1 = 1
			x2 = 1
		}
		t := 1
		for j := i; j < sl && r < sl; {
			result[r] = s[j]
			r++
			if x1*x2 != 0 {
				j = j + x1*t*t + x2*(t-1)*(t-1)
				t = (t - 1) * (t - 1)
			} else {
				j = j + x1 + x2
			}
		}
	}
	re := string(result[:])
	return re
}