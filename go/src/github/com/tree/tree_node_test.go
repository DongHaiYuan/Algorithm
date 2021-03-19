package tree

import "testing"

func TestNewNode(t *testing.T) {
	value := 1
	node := newNode(value)
	if node.value != value {
		t.Errorf("expected: %d, actual: %d", value, node.value)
	}
}

func TestToString(t *testing.T) {
	value := 1
	node := newNode(value)
	expected := "node[1]\n\tleft[nil]\n\tright[nil]"
	if actual := node.String(); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}