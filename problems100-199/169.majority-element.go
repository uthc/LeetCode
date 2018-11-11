/* 169. 求众数
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2 */

func majorityElement(nums []int) int {
	c := make(map[int]int)
	for _, v := range nums {
		if value, ok := c[v]; ok {
			c[v] = value + 1
			if value+1 >= len(nums)/2+1 {
				return v
			}
		} else {
			c[v] = 1
		}
	}
	return nums[0]
}