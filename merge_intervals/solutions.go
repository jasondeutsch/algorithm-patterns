package merge_intervals

import "sort"

type Interval struct {
	Start int
	End   int
}

/*
Given an array of intervals, merge overlapping intervals
input: [[1,4], [7,9],[2,5]]
output: [[1,5], [7,9]]
 */
func MergeSolution(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]Interval, 0, len(intervals))

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= temp.End {
			temp.End = max(temp.End, intervals[i].End)
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)

	return res
}

// helpers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
