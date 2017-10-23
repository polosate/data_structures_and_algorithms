package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(5)
	if !q.isEmpty() {
		t.Error("Expected isEmpty", true, "actual", false)
	}
	q.insert(5)
	q.insert(2)
	q.insert(1)
	q.insert(4)

	el, err := q.remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 1 {
		t.Error("Expected el", 1, "actual", el)
	}

	el, err = q.remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 2 {
		t.Error("Expected el", 2, "actual", el)
	}

	el, err = q.remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 4 {
		t.Error("Expected el", 4, "actual", el)
	}

	el, err = q.remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 5 {
		t.Error("Expected el", 5, "actual", el)
	}
}
