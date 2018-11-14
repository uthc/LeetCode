/* 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2). */

func threeSumClosest(nums []int, target int) int {
	var (
		x int
		y int = nums[0] + nums[1] + nums[2]
	)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				x = nums[i] + nums[j] + nums[k] - target
				if x*x < (y-target)*(y-target) {
					y = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return y
}
