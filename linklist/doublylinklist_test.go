package linklist_test

import (
	"fmt"
	"testing"

	"github.com/puxin71/DesignPatternInGo/linklist"
	"github.com/stretchr/testify/assert"
)

func TestNewDoubleLinkList(t *testing.T) {
	dlist := linklist.NewDoubleLinkList()
	assert.True(t, dlist.IsEmpty(), "created empty double link list")
	assert.NotNil(t, dlist.Top(), "created top sentinel")
	assert.NotNil(t, dlist.Buttom(), "created buttom sentinel")
	assert.Equal(t, 0, dlist.CellCount(), "empty list")
}

func TestAddToButtom(t *testing.T) {
	dlist := linklist.NewDoubleLinkList()
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 9, Prev: nil, Next: nil})
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 8, Prev: nil, Next: nil})
	assert.Equal(t, 3, dlist.CellCount(), "added 3 cells")
	fmt.Printf("Traverse from top: %v\n", dlist.TraverseFromTop())
}

func TestRemoveFromTop(t *testing.T) {
	dlist := linklist.NewDoubleLinkList()
	cell := dlist.RemoveFromTop()
	assert.Nil(t, cell, "return nothing from empty list")
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 9, Prev: nil, Next: nil})
	dlist.AddToButtom(&linklist.DoubleLinkCell{Value: 8, Prev: nil, Next: nil})
	cell = dlist.RemoveFromTop()
	assert.Equal(t, 10, cell.Value.(int), "removed 10 from the list")
	assert.Equal(t, 2, dlist.CellCount(), "left with 2 cells")
	fmt.Printf("Traverse from top: %v\n", dlist.TraverseFromTop())
}
