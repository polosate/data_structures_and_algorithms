package exercises

import "fmt"

type singleList struct {
	first *link
}

func NewSingleList() *singleList {
	return &singleList{
		first: nil,
	}
}

func (sl *singleList) IsEmpty() bool {
	return sl.first == nil
}

func (sl *singleList) DisplayList() {
	current := sl.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println()
}

func (sl *singleList) InsertFirst(value int64) {
	newLink := NewLink(value)
	newLink.next = sl.first
	sl.first = newLink
}

func (sl *singleList) GetIterator() iterator {
	return NewIterator(sl)
}
