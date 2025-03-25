package src

import (
	"testing"
)

/*
TestRootNode_IsRoot tests if the root node claims to be the rootNode
*/
func TestRootNode_IsRoot(t *testing.T) {
	var rootNode = &RootNode{}

	isRoot := rootNode.IsRoot()
	if isRoot != true {
		t.Fatal("Root node does not claim to be root node")
	}
}

/*
TestRootNode_IsLeaf tests if the root node claims to be a leaf
*/
func TestRootNode_IsLeaf(t *testing.T) {
	var rootNode = &RootNode{}

	isLeaf := rootNode.IsLeaf()
	if isLeaf != false {
		t.Fatal("Root node does claim to be leaf node")
	}
}

/*
TestRootNode_ParentIsNil tests if the root node does return nil as parent
*/
func TestRootNode_ParentIsNil(t *testing.T) {
	var rootNode = &RootNode{}

	parent := rootNode.GetParent()
	if parent != nil {
		t.Fatal("Root node should not have parent")
	}
}

/*
TestLeafNode_IsRoot tests if the leaf node claims to be the rootNode
*/
func TestLeafNode_IsRoot(t *testing.T) {
	var leafNode = &LeafNode{}

	isRoot := leafNode.IsRoot()
	if isRoot != false {
		t.Fatal("Leaf node does claim to be root node")
	}
}

/*
TestLeafNode_IsLeaf tests if the leaf node claims to be a leaf
*/
func TestLeafNode_IsLeaf(t *testing.T) {
	var leafNode = &LeafNode{}

	isLeaf := leafNode.IsLeaf()
	if isLeaf != true {
		t.Fatal("Root node does not claim to be leaf node")
	}
}

/*
TestLeafNode_LeftChildIsNil tests if the root node does return nil as parent
*/
func TestLeafNode_LeftChildIsNil(t *testing.T) {
	var leafNode = &LeafNode{}

	child := leafNode.GetLeftChild()
	if child != nil {
		t.Fatal("Leaf node should not have left child")
	}
}

/*
TestLeafNode_RightChildIsNil tests if the root node does return nil as parent
*/
func TestLeafNode_RightChildIsNil(t *testing.T) {
	var leafNode = &LeafNode{}

	child := leafNode.GetRightChild()
	if child != nil {
		t.Fatal("Leaf node should not have right child")
	}
}

/*
TestLeafNode_IsRoot tests if the normal node claims to be the rootNode
*/
func TestNode_IsRoot(t *testing.T) {
	var node = &Node{}

	isRoot := node.IsRoot()
	if isRoot != false {
		t.Fatal("Leaf node does claim to be root node")
	}
}

/*
TestNode_IsLeaf tests if the normal node claims to be a leaf
*/
func TestNode_IsLeaf(t *testing.T) {
	var node = &Node{}

	isLeaf := node.IsLeaf()
	if isLeaf != false {
		t.Fatal("node does claim to be leaf node")
	}
}
