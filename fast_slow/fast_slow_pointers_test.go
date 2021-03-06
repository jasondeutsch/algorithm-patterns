package fast_slow

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	expected := true
	n := &node{0, nil}
	n.next = &node{1, nil}
	n.next.next = &node{2, nil}
	n.next.next.next = &node{3, nil}
	n.next.next.next.next = &node{4, nil}
	n.next.next.next.next.next = &node{5, nil}
	n.next.next.next.next.next.next = n.next.next.next

	if got := DetectCycle(n); got != expected {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}

func TestCycleLength(t *testing.T) {
	expected := 4
	n := &node{0, nil}
	n.next = &node{1, nil}
	n.next.next = &node{2, nil}
	n.next.next.next = &node{3, nil}
	n.next.next.next.next = &node{4, nil}
	n.next.next.next.next.next = &node{5, nil}
	n.next.next.next.next.next.next = n.next.next

	if got := CycleLength(n); got != expected {
		t.Errorf("Error:\n Got: got\n Expected:expected", got, expected)
	}
}
