package entity


// @Created 29/10/2021
// @Updated
type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// @Created 29/10/2021
// @Updated
func (t *TreeNode) Insert(data int) {
	// Binary Search Tree is instantiated but has no head value
	if t.Val == 0 {
		t.Val = data
	} else if data <= t.Val {
		// Input data is less than node value, insert to left side

		// Check left node, if null create left node otherwise insert right away
		if t.Left == nil {
			t.Left = &TreeNode{Val: data}
		} else {
			t.Left.Insert(data)
		}
	} else {
		// Input data is greater than node value, insert to right side

		// Check right node, if null create right node otherwise insert right away
		if t.Right == nil {
			t.Right = &TreeNode{Val: data}
		} else {
			t.Right.Insert(data)
		}
	}
}

// @Created 29/10/2021
// @Updated
func (t *TreeNode) Find(data int) bool {
	searchResult := searchBST(t, data)
	if searchResult == nil {
		return false
	}

	return true
}

// @Created 29/10/2021
// @Updated
func searchBST(root *TreeNode, val int) *TreeNode {

	// Check TreeNode pointer to avoid nil pointer dereference
	if root == nil {
		return nil
	}

	// Input search value is found at the head of the node
	if root.Val == val {
		return root
	}

	// Input search value is less than node value, search left side
	if root.Val > val {
		return searchBST(root.Left, val)
	}

	// Input search value is greater than node value, search right side
	if root.Val < val {
		return searchBST(root.Right, val)
	}

	// Return default value if all conditions are not fulfilled
	return nil
}