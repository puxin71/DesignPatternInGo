package sort_test

import (
	"fmt"
	"testing"

	"github.com/puxin71/DesignPatternInGo/sort"
)

func TestInsertionSort(t *testing.T) {
	list := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.InsertionSort(list)
	fmt.Printf("sorted array: %v\n", list)
}
