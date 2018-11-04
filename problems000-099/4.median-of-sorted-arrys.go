/* 4. 两个排序数组的中位数

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

你可以假设 nums1 和 nums2 不同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

中位数是 (2 + 3)/2 = 2.5 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sumNums []int
	sumLength := len(nums1) + len(nums2)
	var median float64
	n1 := 0
	n2 := 0
	for n1 < len(nums1) && n2 < len(nums2) {
		if nums1[n1] < nums2[n2] {
			sumNums = append(sumNums, nums1[n1])
			n1++
			if n1 == len(nums1) {
				break
			}
		} else {
			sumNums = append(sumNums, nums2[n2])
			n2++
			if n2 == len(nums2) {
				break
			}
		}
	}
	for ; n1 < len(nums1); n1++ {
		sumNums = append(sumNums, nums1[n1])
	}
	for ; n2 < len(nums2); n2++ {
		sumNums = append(sumNums, nums2[n2])
	}
	if sumLength%2 == 0 {
		median = float64(sumNums[sumLength/2-1]+sumNums[sumLength/2]) / 2.0
	} else {
		median = float64(sumNums[sumLength/2])
	}
	return median
}