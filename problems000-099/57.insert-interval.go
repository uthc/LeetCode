/* 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。 */

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	if intervals[0].Start > newInterval.End {
		return append([]Interval{newInterval}, intervals...)
	}
	if intervals[len(intervals)-1].End < newInterval.Start {
		return append(intervals, newInterval)
	}
	insert := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i].Start >= newInterval.Start {
			if intervals[i].Start > newInterval.End {
				return append(intervals[:i], append([]Interval{newInterval}, intervals[i:]...)...)
			}
			intervals[i].Start = newInterval.Start
			if intervals[i].End < newInterval.End {
				intervals[i].End = newInterval.End
			} else {
				return append(intervals[:i], intervals[i:]...)
			}
			insert = i
			break
		}
		if intervals[i].End >= newInterval.Start {
			if intervals[i].End < newInterval.End {
				intervals[i].End = newInterval.End
				if i+1 < len(intervals) && intervals[i].End < intervals[i+1].Start {
					return intervals
				}
			} else {
				return intervals
			}
			insert = i
			break
		}
	}
	if insert+1 < len(intervals) && intervals[insert].End < intervals[insert+1].Start {
		return intervals
	}
	h := intervals[:insert]
	m, l := intervals[insert], []Interval{}
	for i := insert + 1; i < len(intervals); i++ {
		if intervals[i].Start > newInterval.End {
			l = intervals[i:]
			break
		}
	}
	for i := insert + 1; i < len(intervals); i++ {
		if intervals[i].Start > newInterval.End {
			break
		}
		if intervals[i].End > newInterval.End {
			m.End = intervals[i].End
		}
	}

	return append(append(h, m), l...)
}