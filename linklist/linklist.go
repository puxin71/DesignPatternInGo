package linklist

import (
	"fmt"
)

type Element struct {
	Value int
	Next  *Element
}

type linkList struct {
	top    *Element
	tail   *Element
	length int
}

type LinkList interface {
	AddAtBeginning(cell *Element)
	AddAtEnd(cell *Element)
	Delete(value int)
	FindBefore(value int) *Element
	DeleteAll()
	Print()
	Top() *Element
	Tail() *Element
	Length() int
}

func NewLinkList() LinkList {
	// create a sentinel for the new list
	sentinel := Element{Value: 0, Next: nil}
	top, tail := &sentinel, &sentinel
	return &linkList{
		top:    top,
		tail:   tail,
		length: 0,
	}
}

func (l *linkList) AddAtBeginning(cell *Element) {
	if cell == nil {
		return
	}
	cell.Next = l.top.Next
	l.top.Next = cell
	l.length++
	if cell.Next == nil {
		l.tail = cell
	}
}

func (l *linkList) AddAtEnd(cell *Element) {
	if cell == nil {
		return
	}
	l.tail.Next = cell
	l.tail = cell
}

func (l *linkList) Print() {
	var values []int
	iterator := l.top

	for iterator != l.tail.Next {
		if iterator.Next != nil {
			values = append(values, iterator.Next.Value)
		}
		iterator = iterator.Next
	}

	if len(values) == 0 {
		fmt.Println("empty link list")
	}
	fmt.Printf("values in link list: %v", values)
}

func (l *linkList) FindBefore(value int) *Element {
	iterator := l.top
	for iterator != l.tail.Next {
		if iterator.Next.Value == value {
			return iterator
		}
		iterator = iterator.Next
	}
	return iterator
}

func (l *linkList) Delete(value int) {
	iterator := l.top
	for iterator != l.tail.Next {
		if iterator.Next.Value == value {
			iterator.Next = iterator.Next.Next
			l.length--
			return
		}
		iterator = iterator.Next
	}
}

func (l *linkList) Top() *Element {
	return l.top
}

func (l *linkList) Tail() *Element {
	return l.tail
}

func (l *linkList) Length() int {
	return l.length
}

func (l *linkList) DeleteAll() {
	l.top.Next, l.tail.Next = nil, nil
	l.length = 0
}
