// normal binary search tree
// duplicated values are not allowed
package tree

type BST struct {
	root *node
}

// create an empty binary search tree
func NewBST() *BST {
	return &BST{nil}
}

// Insert node to the binary search tree
// return true if succeed to insert the node
// false otherwise
func (bst *BST) Insert(value int) bool {
	var ok bool
	bst.root, ok = insert(bst.root, value)

	return ok
}

func insert(root *node, value int) (*node, bool) {
	if root == nil {
		return newNode(value), true
	}
	if root.value == value {
		return root, false
	} else {
		var ok bool
		if root.value > value {
			root.left, ok = insert(root.left, value)
		} else {
			root.right, ok = insert(root.right, value)
		}

		return root, ok
	}
}

func contains(root *node, value int) bool {
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	} else {
		if root.value > value {
			return contains(root.left, value)
		} else {
			return contains(root.right, value)
		}
	}
}

// Return true if contains the node
// false otherwise
func (bst *BST) Contains(value int) bool {
	return contains(bst.root, value)
}

// Max node in the left subtree
func max(root *node) *node {
	for root.right != nil {
		root = root.right
	}

	return root
}

// delete help function
func delete(root *node, value int) *node {
	if root == nil {
		return root
	}
	if root.value == value {
		// no child
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left != nil && root.right == nil {
			return root.left
		} else if root.left == nil && root.right != nil {
			return root.right
		} else {
			// two children
			left := max(root.left)  // find max value in the left subtree
			root.value = left.value // swap the root value
			root.left = delete(root.left, root.value)

			return root
		}
	} else {
		if root.value > value {
			root.left = delete(root.left, value)
		} else {
			root.right = delete(root.right, value)
		}

		return root
	}
}

// Delete node from the binary search tree
func (bst *BST) Delete(value int) {
	bst.root = delete(bst.root, value)
}




