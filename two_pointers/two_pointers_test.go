package two_pointers

import (
	"testing"
)

func TestTargetSumPairSolution(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := []int{0, 5}
	got := TargetSumPairSolution(nums, target)
	if len(got) != 2 || got[0] != expected[0] || got[1] != expected[1] {
		t.Errorf("Error:\n %v: got\n Expected: %v", got, expected)
	}
}

func TestTargetSumPairs(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := [][]int{{0, 5}, {1, 4}, {2, 3}}
	got := TargetSumPairsSolution(nums, target)
	if len(got) != 3 || got[0][0] != expected[0][0] || got[0][1] != expected[0][1] ||
		got[1][0] != expected[1][0] || got[1][1] != expected[1][1] ||
		got[2][0] != expected[2][0] || got[2][1] != expected[2][1] {
		t.Errorf("Error:\n %v: got\n Expected: %v", got, expected)
	}
}

func TestTargetSumPairHAsh(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	expected := []int{2, 3}
	got := TargetSumPairHash(nums, target)
	if len(got) != 2 || got[0] != expected[0] || got[1] != expected[1] {
		t.Errorf("Error:\n Got: %v\n Expected: %v", got, expected)
	}
}

func TestTargetSumTriplets(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 1, 2, 3}
	target := 0
	expected := [][]int{{-1,3,-2}, {0,2,-2}}
	got := TargetSumTriplets(nums, target)
	for i := range expected[0] {
		if got[0][i] != expected[0][i] ||
			got[1][i] != expected[1][i] {

			t.Errorf("Error:\n Got: %v\n Expected: %v", got, expected)
		}
	}
}
