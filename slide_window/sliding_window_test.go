package slide_window

import (
	"algorithm-patterns/test_helpers"
	"testing"
)


func TestSumSubBrute(t *testing.T) {
	ints := []int{1,2,3,4,5,6,7,8,9,10}
	expected := []int{15, 20, 25, 30, 35, 40}
	got := sumSubBrute(5, ints)
	if !test_helpers.IntSliceEquals(got,expected) {
		t.Errorf(test_helpers.FailString(got, expected))
	}
}

func TestSumSub(t *testing.T) {
	ints := []int{1,2,3,4,5,6,7,8,9,10}
	expected := []int{15, 20, 25, 30, 35, 40}
	got := sumSub(5, ints)
	if !test_helpers.IntSliceEquals(got, expected) {
		t.Errorf(test_helpers.FailString(got, expected))
	}
}

func TestLongestSubStringKDistinct(t *testing.T) {
	str := "aarrbcdeefffe"
	expected := "eefffe"
	got := longestSubStringKDistinct(2, str)
	if got != expected  {
		t.Errorf(test_helpers.FailString(got, expected))
	}

	str = "aarrbcdee"
	expected = "aarr"
	got = longestSubStringKDistinct(2, str)
	if got != expected  {
		t.Errorf(test_helpers.FailString(got, expected))
	}
}

func TestLongestUniqueSubString(t *testing.T) {
	// TODO double check expected
	str := "lakdjflasdjf"
	expected := "kdjflas"
	got := longestUniqueSubString(str)
	if got != expected {
		t.Error(test_helpers.FailString(got, expected))
	}

}
