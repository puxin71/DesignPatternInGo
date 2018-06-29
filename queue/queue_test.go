package queue_test

import (
	"testing"

	"github.com/puxin71/DesignPatternInGo/linklist"

	"github.com/stretchr/testify/assert"

	"github.com/puxin71/DesignPatternInGo/queue"
)

func TestNewQueue(t *testing.T) {
	queue := queue.NewQueue()
	assert.True(t, queue.IsEmpty(), "created empty queue")
	assert.NotNil(t, queue, "allocated memory for queue")
}

func TestEnqueue(t *testing.T) {
	queue := queue.NewQueue()
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 11, Prev: nil, Next: nil})
	assert.Equal(t, 3, queue.Size(), "queued 3 cells")
	queue.PrintFromTop()
}

func TestDequeue(t *testing.T) {
	queue := queue.NewQueue()
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 10, Prev: nil, Next: nil})
	queue.Enqueue(&linklist.DoubleLinkCell{Value: 11, Prev: nil, Next: nil})
	cells := make([]*linklist.DoubleLinkCell, 4)
	for i := 0; i < 4; i++ {
		cells[i] = queue.Dequeue()
	}
	assert.Equal(t, cells[0].Value, cells[1].Value, "dequeued the first two cells")
	assert.Equal(t, 11, cells[2].Value, "dequeued 11 as the third element")
	assert.Nil(t, cells[3], "returned nil out of bound cell")
}
