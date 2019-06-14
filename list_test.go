package list

import (
	"testing"
)

const (
	i1 = "test"
	i2 = 123
	i3 = true
	i4 = "4th"
)

func createList() (l *List) {
	l = NewList()

	l.PushBack(i2)
	l.PushFront(i1)
	l.PushBack(i3)
	// ["test", 123, true]
	return
}

func TestNewListCreation(t *testing.T) {
	l := NewList()
	if l.data == nil {
		t.Error("List has not been created")
	}
	if l.Len() != 0 {
		t.Error("New list has not zero len")
	}
	if l.First() != nil {
		t.Error("New list first element is not nil")
	}
	if l.Last() != nil {
		t.Error("New list last element is not nil")
	}
}

func TestInsertOneItem(t *testing.T) {
	l := NewList()
	l.PushBack(i1)

	if l.First().value != i1 {
		t.Errorf("First item expected %v got %v", i1, l.First().value)
	}
	if l.Last().value != i1 {
		t.Errorf("Last item expected %v got %v", i1, l.Last().value)
	}
}

func TestInsertSeveralItems(t *testing.T) {
	l := createList()
	// check values
	if l.First().value != i1 {
		t.Errorf("First item expected %v got %v", i1, l.First().value)
	}
	if l.Last().value != i3 {
		t.Errorf("Last item expected %v got %v", i3, l.Last().value)
	}

	// check neighbors
	if l.First().Next().value != i2 {
		t.Errorf("First item right neighbor expected %v got %v", i2, l.First().next.value)
	}
	if l.Last().Prev().value != i2 {
		t.Errorf("Last item left neighbor expected %v got %v", i2, l.Last().prev.value)
	}
}

func TestGetNthItem(t *testing.T) {
	l := createList()

	nth2 := l.GetNth(1).Value()
	nth1 := l.GetNth(0).Value()
	nth10 := l.GetNth(10).Value()
	nthMinus2 := l.GetNth(-1).Value()
	nthMinus10 := l.GetNth(-10).Value()

	if nth2 != i2 {
		t.Errorf("Get middle value expected %v got %v", i2, nth2)
	}
	if nth1 != i1 {
		t.Errorf("Get first value expected %v got %v", i1, nth1)
	}
	if nth10 != i3 {
		t.Errorf("Get last value expected %v got %v", i3, nth10)
	}
	if nthMinus2 != i2 {
		t.Errorf("Get middle value expected %v got %v", i2, nthMinus2)
	}
	if nthMinus10 != i1 {
		t.Errorf("Get first value expected %v got %v", i1, nthMinus10)
	}

}

func TestRemoveItem(t *testing.T) {
	l := createList()

	toDelItem := l.GetNth(1)
	toDelItem.Remove()

	if l.First().Next() != l.Last().Prev() {
		t.Error("After remove middle item, first intem does not point to last item")
	}

	l = createList()
	l.First().Remove()
	if l.First().Prev() != nil {
		t.Error("After remove first item, new first item pointer to previos is not nil")
	}
	if l.First().Value() != i2 {
		t.Errorf("After remove first item, new first item expected %v, got %v", i2, l.First().Value())
	}

}
