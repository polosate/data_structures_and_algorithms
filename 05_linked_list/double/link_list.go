package double

import "fmt"

type IDoubleLinkedList interface {
	InsertFirst(dData float32)
	InsertLast(dData float32)
	InsertAfter(key, dData float32)
	DeleteFirst() *link
	DeleteLast() *link
	DeleteAfter(key float32) *link
	DisplayForward()
	DisplayBackward()
	IsEmpty() bool
}

type doubleLinkedList struct {
	first *link
	last  *link
}

func NewDoubleLinkedList() IDoubleLinkedList {
	return &doubleLinkedList{
		first: nil,
		last:  nil,
	}
}

func (ll *doubleLinkedList) IsEmpty() bool {
	return ll.first == nil
}

func (ll *doubleLinkedList) InsertFirst(dData float32) {
	l := NewLink(dData)
	if ll.IsEmpty() {
		ll.last = l
	} else {
		ll.first.previous = l
	}
	l.next = ll.first
	ll.first = l
}

func (ll *doubleLinkedList) InsertLast(dData float32) {
	l := NewLink(dData)
	if ll.IsEmpty() {
		ll.first = l
	} else {
		ll.last.next = l
	}
	l.previous = ll.last
	ll.last = l
}

func (ll *doubleLinkedList) InsertAfter(key, dData float32) {
	l := NewLink(dData)
	if ll.IsEmpty() {
		return
	}
	current := ll.first
	for current != nil && current.dData != key {
		current = current.next
	}
	if current == nil {
		return
	} else if current == ll.last {
		l.next = ll.last.next
		ll.last = l
	} else {
		l.next = current.next
		current.next.previous = l
	}
	current.next = l
	l.previous = current
}

func (ll *doubleLinkedList) DeleteFirst() *link {
	if ll.IsEmpty() {
		return nil
	}
	temp := ll.first
	if temp.next == nil {
		ll.first = nil
		ll.last = nil
	} else {
		temp.next.previous = nil
		ll.first = temp.next
	}
	return temp
}

func (ll *doubleLinkedList) DeleteLast() *link {
	if ll.IsEmpty() {
		return nil
	}
	temp := ll.last
	if temp.previous == nil {
		ll.first = nil
		ll.last = nil
	} else {
		temp.previous.next = nil
		ll.last = temp.previous
	}
	return temp
}

func (ll *doubleLinkedList) DeleteAfter(key float32) *link {
	if ll.IsEmpty() {
		return nil
	}
	current := ll.first
	for current != nil && current.dData != key && current != ll.last {
		current = current.next
	}
	if current == nil || current == ll.last {
		return nil
	} else {
		t := current.next
		current.next = t.next
		current.next.previous = current
		return t
	}
}

func (ll *doubleLinkedList) DisplayForward() {
	fmt.Println("List Forward (first --> last)")
	current := ll.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("\n------")
}

func (ll *doubleLinkedList) DisplayBackward() {
	fmt.Println("List Backward (last --> first)")
	current := ll.last
	for current != nil {
		current.DisplayLink()
		current = current.previous
	}
	fmt.Println("\n------")
}
