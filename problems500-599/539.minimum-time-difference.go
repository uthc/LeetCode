import "strconv"

/* 最小时间差
给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并已分钟数表示。


示例 1：

输入: ["23:59","00:00"]
输出: 1

备注:

列表中时间数在 2~20000 之间。
每个时间取值在 00:00~23:59 之间。 */

import "strconv"

/* 最小时间差
给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并已分钟数表示。


示例 1：

输入: ["23:59","00:00"]
输出: 1

备注:

列表中时间数在 2~20000 之间。
每个时间取值在 00:00~23:59 之间。 */

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	t := func(t string) int {
		h, _ := strconv.Atoi(t[:2])
		m, _ := strconv.Atoi(t[3:])
		return h*60 + m
	}
	f := func(t1, t2 int) int {
		dt := t1 - t2
		if dt < 0 {
			dt = -dt
		}
		if dt > 720 {
			return 1440 - dt
		}
		return dt
	}
	minutes := make([]int, len(timePoints))
	for i, v := range timePoints {
		minutes[i] = t(v)
	}
	minutes = Sort(minutes)
	min := f(minutes[0], minutes[len(minutes)-1])
	for i := 1; i < len(minutes); /* && min != 0 */ i++ {
		dt := f(minutes[i], minutes[i-1])
		if dt < min {
			min = dt
		}
	}
	return min
}

func Sort(t []int) []int {
	r := make([]int, 0)
	if len(t) < 2 {
		return append(r, t...)
	}
	t1, t2 := Sort(t[:len(t)/2]), Sort(t[len(t)/2:])
	i, j := 0, 0
	for i < len(t1) && j < len(t2) {
		if t1[i] < t2[j] {
			r = append(r, t1[i])
			i++
		} else {
			r = append(r, t2[j])
			j++
		}
	}
	r = append(r, t1[i:]...)
	r = append(r, t2[j:]...)
	return r
}