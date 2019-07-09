package merge_intervals

import (
	"testing"
)

func TestMergeSolution(t *testing.T) {
	intervals := []Interval{
		{1, 6},
		{2, 7},
		{8, 9},
	}

	expected := []Interval{
		{1, 7},
		{8, 9},
	}

	got := MergeSolution(intervals)
	for i := range got {
		if got[i].Start != expected[i].Start || got[i].End != expected[i].End {
			t.Errorf("Error:\n Got: %v\n Expected: %v", got, expected)
		}
	}
}

func TestMaxDisjointIntervals(t *testing.T) {
	// input: [[1,4], [7,11],[7,9],[9,10],[2,3], [4,5]]
	// output: [[2,3],[4,5],[7,9]]
	input := []Interval{
		{1, 4},
		{7, 11},
		{7, 9},
		{9, 10},
		{2, 3},
		{4, 5},
	}
	expected := []Interval{
		{2, 3},
		{4, 5},
		{7, 9},
		{9, 10},
	}

	got := MaxDisjointIntervals(input)
	for i := range got {
		if got[i].Start != expected[i].Start ||
			got[i].End != expected[i].End {
			t.Errorf("Error:\n Got: %v\n Expected: %v", got, expected)
		}
	}
}
