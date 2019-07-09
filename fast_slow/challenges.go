package fast_slow

// DEMO CODE

type node struct {
	val  int
	next *node
}

/*******************************

Does a given linked list contain a cycle?

*******************************/
func DetectCycle(n *node) bool {
	slow, fast := n, n

	for  fast != nil && fast.next != nil {
		if fast == slow{
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return false
}

/*******************************

Find the length of the cycle.
return 0 if no cycle exists.

*******************************/
func CycleLength(head *node) int {
	return 0
}
