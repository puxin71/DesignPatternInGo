package balancedtree

import "fmt"

// go provides a good heap library
type heap struct {
	values []int
}

type Heap interface {
	Values() []int
	PrintHeapTree()
	Add(value int)

	// removeFromTop (rebalance heap takes O(log(N)ß))
}

// heap stores complete binary tree in an array. It takes O(Nlog(N)) steps
func NewHeap(values []int) Heap {
	if len(values) == 0 {
		return nil
	}
	fmt.Printf("input array: %v\n", values)

	for i := 0; i < len(values); i++ {
		index := i
		for index != 0 {
			parentIndex := int((index - 1) / 2)
			if values[index] <= values[parentIndex] {
				break
			}

			fmt.Printf("swapping parentIdx: %d, value:%d with idx: %d, value: %d\n",
				parentIndex, values[parentIndex], index, values[index])
			values[parentIndex], values[index] = values[index], values[parentIndex]

			// move to the parent index
			index = parentIndex
		}
	}
	return &heap{values: values}
}

func (h *heap) Values() []int {
	return h.values
}

func (h *heap) PrintHeapTree() {
	for i := 0; i < len(h.values); i++ {
		fmt.Printf("node: %d", h.values[i])
		if (i*2 + 1) < len(h.values) {
			fmt.Printf(",\tleft child: %d", h.values[i*2+1])
		}
		if (i*2 + 2) < len(h.values) {
			fmt.Printf(",\tright child: %d", h.values[i*2+2])
		}
		fmt.Printf("\n")
	}
}

// Add a new item takes O(log(N)), which is the depth of a heap binary tree
func (h *heap) Add(value int) {
	h.values = append(h.values, value)
	nHeap := NewHeap(h.values)
	h.values = nHeap.Values()
}
