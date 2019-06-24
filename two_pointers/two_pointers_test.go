package two_pointers

import (
	"study/test_helpers"
	"testing"
)

func TestTargetSumPair(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := []int{0, 5}
	got := TargetSumPair(nums, target)
	if len(got) != 2 || got[0] != expected[0] || got[1] != expected[1] {
		t.Error(test_helpers.FailString(got, expected))
	}
}

func TestTargetSumPairs(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := [][]int{{0, 5}, {1, 4}, {2, 3}}
	got := TargetSumPairs(nums, target)
	if len(got) != 3 || got[0][0] != expected[0][0] || got[0][1] != expected[0][1] ||
		got[1][0] != expected[1][0] || got[1][1] != expected[1][1] ||
		got[2][0] != expected[2][0] || got[2][1] != expected[2][1] {
		t.Error(test_helpers.FailString(got, expected))
	}
}

func TestTargetSumPairHAsh(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := []int{2, 3}
	got := TargetSumPairHash(nums, target)
	if len(got) != 2 || got[0] != expected[0] || got[1] != expected[1] {
		t.Error(test_helpers.FailString(got, expected))
	}
}

func TestTargetSumTriplets(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 1, 2, 3}
	target := 0
	expected := [][]int{}
	got := TargetSumTriplets(nums, target)
	if true {
		t.Error(test_helpers.FailString(got, expected))
	}
}
