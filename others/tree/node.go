package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n Node) Print() {
	fmt.Print(n.Value, " ")
}

func (n *Node) SetValue(value int) {
	if n == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	n.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
