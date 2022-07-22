package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var root Node

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("Node Count:", nodeCount)
}
