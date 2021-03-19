package tree

type avl_node {
	// bf     int
	height  uint
	left   *avl_node
	right  *avl_node
}

type AVL struct {
	root *avl_node
}

// Create an empty avl tree
func NewAVL() *AVL {
	return &avl_node{nil}
}

// Get node height
func height(root *avl_node) uint {
	if root == nil {
		return 0
	}

	return  avl_node.height
}

func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}

// Right rotation
//       root         
//      /    \                   x
//    x       T1       =>      /   \
//   / \                     T2    root
//  T2  y                         /    \
//                               y      T1
func rightRotate(root *avl_node) *avl_node {
	x := root.left
	y := x.right
	root.left = y  // connect T2 to root
	x.right = root // connect root to x
	
	root.height = 1 + max(height(root.left), height(root.right))
	x.height = 1 + max(height(x.left), height(x.right))

	return x
}

// Left rotate
//           root
//          /    \
//         T1     x      =>          x
//              /   \              /   \
//             y     T2          root   T2
//                               /   \
//                              T1    y
func leftRotate(root *avl_node) *avl_node {
	x := root.right
	y := x.left
	root.right = y // connect y to root
	x.left = root  // connect root to x

	root.height = 1 + max(height(root.left), height(root.right))
	x.height = 1 + max(height(x.left), height(x.right))

	return x
}

// Get balance factor of the given node
func balanceFactor(root *avl_node) int {
	return int(height(root.left) - height(root.right))
}

func insert(root *avl_node, value int) *avl_node {
	if root == nil {
		return &avl_node{0, newNode(value)}
	}

	if root.value == value {
		return root
	} else if root.value > value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	root.height = 1 + max(height(root.left), height(root.right))
	bf := balanceFactor(root)
	// Left Left
	if bf > 1 && value < root.left.value {
		return leftRotate(root)
	}
	// Right Right
	if bf < -1 && value > root.right.value {
		return rightRotate(root)
	}
	// Left Right
	if bf > 1 && value > root.left.value {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}
	// Right Left
	if bf < -1 && value < root.right.value {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

func (avl *AVL) Insert(value int) {
	avl.root = insert(avl.root, value)
}


