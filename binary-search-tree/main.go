package bst

type Node struct {
	Val   int
	Count int
	Left  *Node
	Right *Node
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}

	if root.Val == val {
		root.Count++
	}
	if root.Val < val {
		root.Right = Insert(root.Right, val)
	} else {
		root.Left = Insert(root.Left, val)
	}

	return root
}
