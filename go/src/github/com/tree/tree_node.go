package tree

import (
	"bytes"
	"fmt"
)

type node struct {
	value  int
	left  *node
	right *node
}

// Create a new node without children
func newNode(value int) *node {
	return &node{value, nil, nil}
}

func (root node) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("node[%d]", root.value))
	if root.left != nil {
		buf.WriteString(fmt.Sprintf("\n\tleft[%d]", root.left.value))
	} else {
		buf.WriteString("\n\tleft[nil]")
	}
	if root.right != nil {
		buf.WriteString(fmt.Sprintf("\n\tright[%d]", root.right.value))
	} else {
		buf.WriteString("\n\tright[nil]")
	}

	return buf.String()
}