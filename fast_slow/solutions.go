package fast_slow

// classic problem for this pattern
func detectCycleSolution(n *node) bool {
	fast, slow := n, n
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func cycleLengthSolution(head *node) int {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		// first detect cycle
		if slow == fast {
			count := 1
			slow = slow.next
			for fast != slow {
				count++
				slow = slow.next
			}
			return count
		}
	}
	return 0
}
