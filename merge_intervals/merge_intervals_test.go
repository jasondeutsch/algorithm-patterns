package merge_intervals

import (
	"testing"
)

func TestMergeSolution(t *testing.T) {
	intervals := []Interval{
		Interval{1,6},
		Interval{2,7},
		Interval{8,9},
	}

	expected := []Interval{
		Interval{1,7},
		Interval{8,9},
	}

	got :=  MergeSolution(intervals)
	for i := range got {
		if got[i].Start != expected[i].Start || got[i].End != expected[i].End {
			t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
		}
	}


}
