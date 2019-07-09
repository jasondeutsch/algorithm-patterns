package sliding_window

import (
	"testing"
)

func TestSumSubBrute(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{15, 20, 25, 30, 35, 40}
	got := sumSubBrute(5, ints)
	if !IntSliceEquals(got, expected) {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}

func TestSumSubSolution(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{15, 20, 25, 30, 35, 40}
	got := sumSubSolution(5, ints)
	if !IntSliceEquals(got, expected) {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}

func TestLongestSubStringKDistinct(t *testing.T) {
	str := "aarrbcdeefffe"
	expected := "eefffe"
	got := longestSubStringKDistinct(2, str)
	if got != expected {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}

	str = "aarrbcdee"
	expected = "aarr"
	got = longestSubStringKDistinct(2, str)
	if got != expected {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}

func TestLongestUniqueSubString(t *testing.T) {
	// TODO double check expected
	str := "lakdjflasdjf"
	expected := "kdjflas"
	got := longestUniqueSubStringSolution(str)
	if got != expected {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}

func IntSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
