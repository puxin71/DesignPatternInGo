package linklist_test

import (
	"testing"

	"github.com/puxin71/DesignPatternInGo/linklist"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkList(t *testing.T) {
	list := linklist.NewLinkList()
	assert.NotNil(t, list, "create not null link list with a sentinel element")
}

// assuming that the list allows duplicated values
func TestAddAtBeginning(t *testing.T) {
	list := linklist.NewLinkList()
	cell := &linklist.Element{Value: 1, Next: nil}
	list.AddAtBeginning(cell)
	top := list.Top()
	assert.Equal(t, cell.Value, top.Next.Value, "added at beginning")
	cell = nil
	list.AddAtBeginning(cell)
	top = list.Top()
	assert.Equal(t, 1, top.Next.Value, "did not add empty cell")
	cell = &linklist.Element{Value: 10, Next: nil}
	list.AddAtBeginning(cell)
	top = list.Top()
	assert.Equal(t, 10, top.Next.Value, "added 10 at beginning")
}

func TestAddAtEnd(t *testing.T) {
	list := linklist.NewLinkList()
	cell := &linklist.Element{Value: 1, Next: nil}
	list.AddAtBeginning(cell)
	cell = &linklist.Element{Value: 10, Next: nil}
	list.AddAtEnd(cell)
	tail := list.Tail()
	assert.Equal(t, 10, tail.Value, "added 10 at the end")
	cell = nil
	list.AddAtEnd(cell)
	tail = list.Tail()
	assert.Equal(t, 10, tail.Value, "did not add empty cell")
}

func TestFindBefore(t *testing.T) {

}
