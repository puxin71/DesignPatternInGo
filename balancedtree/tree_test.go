package balancedtree_test

import (
	"fmt"
	"testing"

	"github.com/puxin71/DesignPatternInGo/balancedtree"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeAdd(t *testing.T) {
	btree := balancedtree.NewBinaryTree()
	btree.Add(&balancedtree.Node{Value: 10, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 1, btree.NodeCount(), "added one node")
	btree.Add(&balancedtree.Node{Value: 8, LeftChild: nil, RightChild: nil})
	btree.Add(&balancedtree.Node{Value: 11, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 3, btree.NodeCount(), "added 3 nodes")
	btree.Add(&balancedtree.Node{Value: 4, LeftChild: nil, RightChild: nil})
	btree.Add(&balancedtree.Node{Value: 15, LeftChild: nil, RightChild: nil})
	btree.Add(&balancedtree.Node{Value: 7, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 6, btree.NodeCount(), "added 6 nodes")

	// print values in preOrdered traversal
	values := btree.PreOrderTraversal()
	fmt.Printf("preOrder traversed values: %v\n", values)
}
