package linklist

import (
	"math"
)

type DoubleLinkCell struct {
	Value int
	Prev  *DoubleLinkCell
	Next  *DoubleLinkCell
}

type doubleLinkList struct {
	topSentinel    *DoubleLinkCell
	bottomSentinel *DoubleLinkCell
}

type DoubleLinkList interface {
	Top() *DoubleLinkCell
	Buttom() *DoubleLinkCell
	IsEmpty() bool
	RemoveFromTop() *DoubleLinkCell
	AddToButtom(cell *DoubleLinkCell)
	CellCount() int
	TraverseFromTop() []int
}

func NewSentinel() *DoubleLinkCell {
	return &DoubleLinkCell{
		Value: math.MaxInt64,
		Prev:  nil,
		Next:  nil,
	}
}

func NewDoubleLinkList() DoubleLinkList {
	top := NewSentinel()
	buttom := NewSentinel()
	top.Next = buttom
	buttom.Prev = top

	return &doubleLinkList{
		topSentinel:    top,
		bottomSentinel: buttom,
	}
}

func (db *doubleLinkList) Top() *DoubleLinkCell {
	return db.topSentinel
}

func (db *doubleLinkList) Buttom() *DoubleLinkCell {
	return db.bottomSentinel
}

func (db *doubleLinkList) IsEmpty() bool {
	return (db.topSentinel.Next == db.bottomSentinel &&
		db.bottomSentinel.Prev == db.topSentinel)
}

func (db *doubleLinkList) RemoveFromTop() *DoubleLinkCell {
	if db.IsEmpty() {
		return nil
	}
	cell := db.Top().Next
	db.Top().Next = cell.Next
	cell.Next.Prev = db.Top()
	cell.Prev = nil
	cell.Prev = nil
	return cell
}

func (db *doubleLinkList) AddToButtom(cell *DoubleLinkCell) {
	if cell == nil {
		return
	}
	buttom := db.Buttom()
	prev := buttom.Prev
	prev.Next = cell
	cell.Prev = prev
	cell.Next = buttom
	buttom.Prev = cell
}

func (db *doubleLinkList) CellCount() int {
	count := 0
	iterator := db.Top().Next
	for iterator != db.Buttom() {
		count++
		iterator = iterator.Next
	}
	return count
}

func (db *doubleLinkList) TraverseFromTop() []int {
	var list []int
	iterator := db.Top().Next
	for iterator != db.Buttom() {
		list = append(list, iterator.Value)
		iterator = iterator.Next
	}
	return list
}
