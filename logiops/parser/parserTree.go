package parser

type Node struct {
	Children []Node
	Value string
}

func (n *Node) Append(newNode Node) {
	n.Children = append(n.Children, newNode)
}

func NewNode(val string) Node {
	return Node{Value: val}
}