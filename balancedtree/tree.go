package balancedtree

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

type binaryTree struct {
	root      *Node
	height    int
	nodeCount int
}

type BinaryTree interface {
	Add(node *Node)
	NodeCount() int
	PreOrderTraversal() []int
}

func NewBinaryTree() BinaryTree {
	return &binaryTree{root: nil, height: 0, nodeCount: 0}
}

func (t *binaryTree) Add(node *Node) {
	t.add(t.root, node)
}

func (t *binaryTree) add(parent, node *Node) {
	if node == nil {
		return
	}
	if t.root == nil {
		t.root = node
		t.nodeCount++
		return
	}

	if node.Value == parent.Value {
		return
	}

	if node.Value < parent.Value && parent.LeftChild == nil {
		parent.LeftChild = node
		t.nodeCount++
		return
	}

	if node.Value > parent.Value && parent.RightChild == nil {
		parent.RightChild = node
		t.nodeCount++
		return
	}

	if parent.LeftChild != nil && node.Value < parent.Value {
		t.add(parent.LeftChild, node)
	}

	if parent.RightChild != nil && node.Value > parent.Value {
		t.add(parent.RightChild, node)
	}
}

func (t *binaryTree) NodeCount() int {
	return t.nodeCount
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
