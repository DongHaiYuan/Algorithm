import "testing"

func balanced(root *avl_node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && root.value 

	bf := balanceFactor(root)
	if bf >= 2 || bf <= -2 {
		return false
	}

	left := balanced(root.left)
	right := balanced(root.right)

	return left && right
}


func TestInsert(t *testing.T) {
	avl := NewAVL()
	times := 100
	for i := 0; i < times; i++ {
		avl.Insert(i)
	}

	if 
} 