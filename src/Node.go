package src

type INode interface {
	IsRoot() bool
	IsLeaf() bool
	GetParent() INode
	GetLeftChild() INode
	GetRightChild() INode
}

/*
////// RootNode //////
*/

/*
RootNode object definition
*/
type RootNode struct {
	leftChild  INode
	rightChild INode
}

func (n *RootNode) IsRoot() bool {
	return true
}

func (n *RootNode) IsLeaf() bool {
	return false
}

func (n *RootNode) GetParent() INode {
	return nil
}

func (n *RootNode) GetLeftChild() INode {
	return n.leftChild
}

func (n *RootNode) GetRightChild() INode {
	return n.rightChild
}

/*
////// LeafNode //////
*/

type LeafNode struct {
	parent INode
}

func (n *LeafNode) IsRoot() bool {
	return false
}

func (n *LeafNode) IsLeaf() bool {
	return true
}

func (n *LeafNode) GetParent() INode {
	return n.parent
}

func (n *LeafNode) GetLeftChild() INode {
	return nil
}

func (n *LeafNode) GetRightChild() INode {
	return nil
}

/*
////// Node //////
*/

/*
Node is the representation of a normal node in the middle
*/
type Node struct {
	parent     INode
	leftChild  INode
	rightChild INode
}

func (n *Node) IsRoot() bool {
	return false
}

func (n *Node) IsLeaf() bool {
	return false
}

func (n *Node) GetParent() INode {
	return n.parent
}

func (n *Node) GetLeftChild() INode {
	return n.leftChild
}

func (n *Node) GetRightChild() INode {
	return n.rightChild
}
