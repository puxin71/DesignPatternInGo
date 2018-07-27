package balancedtree_test

import (
	"fmt"
	"testing"

	"github.com/puxin71/DesignPatternInGo/balancedtree"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	btree := balancedtree.NewBinaryTree()
	btree.BinarySortedAdd(&balancedtree.Node{Value: 10, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 1, btree.NodeCount(), "added one node")
	btree.BinarySortedAdd(&balancedtree.Node{Value: 8, LeftChild: nil, RightChild: nil})
	btree.BinarySortedAdd(&balancedtree.Node{Value: 11, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 3, btree.NodeCount(), "added 3 nodes")
	btree.BinarySortedAdd(&balancedtree.Node{Value: 4, LeftChild: nil, RightChild: nil})
	btree.BinarySortedAdd(&balancedtree.Node{Value: 15, LeftChild: nil, RightChild: nil})
	btree.BinarySortedAdd(&balancedtree.Node{Value: 7, LeftChild: nil, RightChild: nil})
	btree.BinarySortedAdd(&balancedtree.Node{Value: 9, LeftChild: nil, RightChild: nil})
	btree.BinarySortedAdd(&balancedtree.Node{Value: 12, LeftChild: nil, RightChild: nil})
	assert.Equal(t, 8, btree.NodeCount(), "added 6 nodes")
	assert.Equal(t, 3, btree.Height(), "added 3 tree levels")

	// print values in preOrdered traversal
	values := btree.PreOrderTraversal()
	fmt.Printf("preOrder traversed values: %v\n", values)

	// print values in inOrdered traversal (sorted)
	values = btree.InOrderTraversal()
	fmt.Printf("inOrder traversed values: %v\n", values)

	// print values in inOrdered traversal
	values = btree.PostOrderTraversal()
	fmt.Printf("postOrder traversed values: %v\n", values)

	// print values in inOrdered traversal
	values = btree.TraverseDepthFirst(btree.Root())
	fmt.Printf("depth first traversed values: %v\n", values)

	node := btree.Get(10)
	assert.Equal(t, btree.Root(), node, "find root value")
	node = btree.Get(12)
	assert.Equal(t, 12, node.Value, "find value 12")
	node = btree.Get(8)
	assert.Equal(t, 8, node.Value, "find value 12")
	node = btree.Get(100)
	assert.Nil(t, node, "100 is not in the tree")
}
