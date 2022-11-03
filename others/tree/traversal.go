package tree

import "fmt"

// func (n *Node) Traverse() {
// 	if n == nil {
// 		return
// 	}

// 	n.Left.Traverse()
// 	n.Print()
// 	n.Right.Traverse()
// }

func (n *Node) Traverse() {
	n.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (n *Node) TraverseFunc(f func(*Node)) {
	if n == nil {
		return
	}

	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
