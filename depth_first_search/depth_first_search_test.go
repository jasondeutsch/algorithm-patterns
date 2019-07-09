package depth_first_search

import "testing"


var testTree = Node {
	Val: 5,
	Left: &Node{Val: 3, Left: nil, Right: nil},
	Right: &Node{Val: 7, Left: nil, Right: nil},
}


func TestDFS(t *testing.T) {
	expected := 7
	if got := DFS(&testTree, 7).Val; got != expected {
		t.Errorf("Error:\n Got: %v\n Expected: %v", got, expected)
	}
}
