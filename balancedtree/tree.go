package balancedtree

import (
	"fmt"
	"math"

	"github.com/puxin71/DesignPatternInGo/linklist"
	"github.com/puxin71/DesignPatternInGo/queue"
)

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

type binaryTree struct {
	root      *Node
	nodeCount int
	height    int
}

type BinaryTree interface {
	Add(node *Node)
	NodeCount() int
	Height() int
	PreOrderTraversal() []int
	InOrderTraversal() []int
	PostOrderTraversal() []int
	TraverseDepthFirst(node *Node) []int
	Root() *Node
	Get(value int) *Node

	// delete node from sorted tree is complex!!!
}

func NewBinaryTree() BinaryTree {
	return &binaryTree{root: nil, nodeCount: 0, height: 0}
}

func (t *binaryTree) Add(node *Node) {
	if t.root == nil {
		t.root = node
		t.nodeCount++
		return
	}
	t.add(t.root, node)
}

// sorted add. left child < node, right child >  node
func (t *binaryTree) add(parent, node *Node) {
	if node == nil {
		return
	}

	if node.Value == parent.Value {
		return
	}

	if node.Value < parent.Value {
		if parent.LeftChild == nil {
			parent.LeftChild = node
			t.nodeCount++
			if parent.RightChild == nil {
				t.height++
			}
			fmt.Printf("parent: %d, add left child: %d\n", parent.Value, parent.LeftChild.Value)
		} else {
			t.add(parent.LeftChild, node)
		}
	} else {
		if parent.RightChild == nil {
			parent.RightChild = node
			t.nodeCount++
			fmt.Printf("parent: %d, add right child: %d\n", parent.Value, parent.RightChild.Value)
		} else {
			t.add(parent.RightChild, node)
		}
	}
}

func (t *binaryTree) NodeCount() int {
	return t.nodeCount
}

func (t *binaryTree) Height() int {
	return int(math.Log2(float64(t.nodeCount)))
}

func (t *binaryTree) PreOrderTraversal() []int {
	var values []int
	return t.traversePreOrder(t.root, values)
}

// the algorithm processes a node, and then its left child, and then its right child
func (t *binaryTree) traversePreOrder(node *Node, values []int) []int {

	values = append(values, node.Value)

	if node.LeftChild != nil {
		values = t.traversePreOrder(node.LeftChild, values)
	}

	if node.RightChild != nil {
		values = t.traversePreOrder(node.RightChild, values)
	}

	return values
}

func (t *binaryTree) InOrderTraversal() []int {
	var values []int
	if t.root == nil {
		return nil
	}
	values = t.traverseInOrder(t.root, values)
	return values
}

// inOrder process the left child, the node, then the right child
func (t *binaryTree) traverseInOrder(node *Node, values []int) []int {
	if node.LeftChild != nil {
		values = t.traverseInOrder(node.LeftChild, values)
	}
	values = append(values, node.Value)
	if node.RightChild != nil {
		values = t.traverseInOrder(node.RightChild, values)
	}
	return values
}

// postOrder traverse the right child, the left child, then the node
func (t *binaryTree) PostOrderTraversal() []int {
	var values []int
	if t.root == nil {
		return nil
	}
	values = t.traversePostOrder(t.root, values)
	return values
}

func (t *binaryTree) traversePostOrder(node *Node, values []int) []int {
	if node.RightChild != nil {
		values = t.traversePostOrder(node.RightChild, values)
	}
	if node.LeftChild != nil {
		values = t.traversePostOrder(node.LeftChild, values)
	}
	values = append(values, node.Value)
	return values
}

func (t *binaryTree) TraverseDepthFirst(node *Node) []int {
	var values []int
	if node == nil {
		return nil
	}
	myQueue := queue.NewQueue()
	myQueue.Enqueue(t.createCell(node))

	for !myQueue.IsEmpty() {
		cell := myQueue.Dequeue()
		currNode := cell.Value.(Node)
		values = append(values, currNode.Value)
		if currNode.LeftChild != nil {
			myQueue.Enqueue(t.createCell(currNode.LeftChild))
		}
		if currNode.RightChild != nil {
			myQueue.Enqueue(t.createCell(currNode.RightChild))
		}
	}
	return values
}

func (t *binaryTree) createCell(node *Node) *linklist.DoubleLinkCell {
	if node == nil {
		return nil
	}
	return &linklist.DoubleLinkCell{Value: *node, Prev: nil, Next: nil}
}

func (t *binaryTree) Root() *Node {
	return t.root
}

func (t *binaryTree) Get(value int) *Node {
	if t.root == nil {
		return nil
	}
	return t.get(t.root, value)
}

func (t *binaryTree) get(parent *Node, value int) *Node {
	if parent.Value == value {
		return parent
	}
	if value < parent.Value {
		if parent.LeftChild == nil {
			return nil
		} else {
			return t.get(parent.LeftChild, value)
		}
	}
	if parent.RightChild == nil {
		return nil
	}
	return t.get(parent.RightChild, value)
}
