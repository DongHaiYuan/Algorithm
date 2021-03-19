package tree

import (
	"fmt"
	"testing"
	"math/rand"
)

func balanced(root *node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && root.value <= root.left.value {
		return false
	}
	if root.right != nil && root.value >= root.right.value {
		return false
	}
	left := balanced(root.left)
	right := balanced(root.right)

	return left && right
}


func TestInsert(t *testing.T) {
	bst := NewBST()
	if ok := bst.Insert(10); !ok {
		t.Errorf("`bst.Insert(10)`, expected: true, actual: false")
	} 

	if ok := bst.Insert(10); ok {
		t.Errorf("`bst.Insert(10)`, expected: false, actual: true")
	} 	

	times := 1000000
	for i := 0; i < times; i++ {
		bst.Insert(rand.Int())
	}
	if !balanced(bst.root) {
		t.Errorf("expected: true, actual: false")
	}
}

func TestDelete(t *testing.T) {
	value := 10
	bst := NewBST()
	bst.Insert(value)
	bst.Delete(value)
	if bst.Contains(value) {
		t.Errorf("expected: false, actual: true")
	}
	times := 10000000
	mp := make(map[int]bool)
	for i := 0; i < times; i++ {
		mp[i] = true
	}
	
	for k, _ := range mp {
		bst.Insert(k)
	}

	for k, _ := range mp {
		bst.Delete(k)
	}

	if !balanced(bst.root) {
		t.Errorf("expected: true, actual: false")
	}

	if bst.root != nil {
		t.Errorf(fmt.Sprintf("expected: nil, actual: %s", bst.root.String()))
	}
}