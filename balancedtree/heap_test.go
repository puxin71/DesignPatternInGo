package balancedtree_test

import (
	"fmt"
	"testing"

	"github.com/puxin71/DesignPatternInGo/balancedtree"
)

func TestNewHeap(t *testing.T) {
	myHeap := balancedtree.NewHeap([]int{10, 8, 11, 4, 9, 15, 7, 12})
	fmt.Printf("heap array: %v\n", myHeap.Values())
	myHeap.PrintHeapTree()

	myHeap.Add(20)
	fmt.Printf("heap array: %v\n", myHeap.Values())
	myHeap.PrintHeapTree()
}
