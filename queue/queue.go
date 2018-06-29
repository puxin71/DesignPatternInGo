package queue

import (
	"fmt"

	"github.com/puxin71/DesignPatternInGo/linklist"
)

type queue struct {
	dlist linklist.DoubleLinkList
}

type Queue interface {
	Enqueue(cell *linklist.DoubleLinkCell)
	Dequeue() *linklist.DoubleLinkCell
	IsEmpty() bool
	Size() int
	PrintFromTop()
}

func NewQueue() Queue {
	return &queue{dlist: linklist.NewDoubleLinkList()}
}

func (q *queue) Enqueue(cell *linklist.DoubleLinkCell) {
	q.dlist.AddToButtom(cell)
}

func (q *queue) Dequeue() *linklist.DoubleLinkCell {
	return q.dlist.RemoveFromTop()
}

func (q *queue) IsEmpty() bool {
	return q.dlist.IsEmpty()
}

func (q *queue) Size() int {
	return q.dlist.CellCount()
}

func (q *queue) PrintFromTop() {
	values := q.dlist.TraverseFromTop()
	fmt.Printf("values from top: %v\n", values)
}
