package depth_first_search

type Node struct {
	Val int
	Left *Node
	Right *Node
}

/*
Find a node of val k
*/

func DFS(node *Node, k int) *Node{
	if node == nil || node.Val == k{
		return node
	} else {
		DFS(node.Left, k)
		DFS(node.Right, k)
	}
	return nil
}

/*
Given the root of a binary search tree with distinct values,
modify it so that every node has a new value equal to the
sum of the values of the original tree that are greater than or equal to node.val.
*/


