package merge_intervals

import (
	"github.com/golang-collections/collections/stack"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

/*
Given an array of intervals, merge overlapping intervals
input: [[1,3], [7,9],[2,5]]
output: [[1,5], [7,9]]
*/

func mergeSolution(intervals []Interval) []Interval {
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

func MergeSolution(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort in increasing order of start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	stack := stack.New()
	stack.Push(intervals[0])
	for _, interval := range intervals[1:] {
		top := stack.Peek().(Interval)

		// compare to top for overlap
		// does not overlap
		if interval.Start > top.End {
			stack.Push(interval)
		} else if interval.End > top.End {
			newTop := stack.Pop().(Interval)
			newTop.End = max(newTop.End, interval.End)
			stack.Push(newTop)
		}
		// hidden else case: a completely overlaps b
	}

	// Stack -> Interval Slice
	var result []Interval
	for {
		if popped, ok := stack.Pop().(Interval); ok {
			result = append(result, popped)
		} else {
			break
		}
	}

	return result
}

/*
Given an array of intervals, select the the most non-pverlapping intervals??
input: [[1,4], [7,11],[7,9],[9,10],[2,3], [4,5]]
output: [[2,3],[4,5],[7,9]]

aka. max disjoint intervals
*/

func MaxDisjointIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].End < intervals[j].End
	})

	var result []Interval
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= intervals[i-1].End {
			result = append(result, intervals[i])
		}
	}

	return result
}

/*

 A conference room is available to be booked between 900 and 1800
find the largest available space of time given an array of bookings.
input: [[900, 1100],[1200,1300],[1500,1600]]
output: [1300,1500]
```
*/

func largestOpenInterval(intervals []Interval) Interval {
	result := Interval{}
	if len(intervals) <= 1 {
		return result
	}

	maxTime := 0
	for i := 1; i < len(intervals); i++ {
		time := intervals[i].Start - intervals[i-1].End
		if time > maxTime {
			maxTime = time
			result.Start = intervals[i-1].End
			result.End = intervals[i].Start
		}
	}

	return result
}

// helpers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
