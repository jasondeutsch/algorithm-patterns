package merge_intervals

import (
	"sort"
)

/*******************************

Merge all overlapping intervals.
Should return an array of arrays.

EX.

Input: [[1,6], [1,4], [2,5], [5,7],[8,9]]
Output: [[1,4][5,7],[8,9]]

*******************************/


func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// step 1: sort intervals by
	sort.Slice(intervals, func(i , j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		if top := result[len(result)-1]; overlaps(top, interval)  { // does it overlap?
		  if top[1] < interval[1] { // does b and after a ?
		  	result[len(result)-1][1] = interval[1]
		  }
		} else {
			result = append(result, interval)
		}
	}

	return result
}

func overlaps(a, b []int) bool {
	return a[1] < b[0] ||
		a[1] < b[1]
}

