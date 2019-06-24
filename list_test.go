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

	nth2 := l.GetNth(1)
	nth1 := l.GetNth(0)

	nthMinus2 := l.GetNth(-1)

	if nth2.Value() != i2 {
		t.Errorf("Get middle value expected %v got %v", i2, nth2)
	}
	if nth1.Value() != i1 {
		t.Errorf("Get first value expected %v got %v", i1, nth1)
	}

	if nthMinus2.Value() != i2 {
		t.Errorf("Get middle value expected %v got %v", i2, nthMinus2)
	}

}

func TestGetNthOutOfBounds(t *testing.T) {
	l := createList()
	nth10 := l.GetNth(10)       // nil
	nthMinus10 := l.GetNth(-10) // nil

	if nth10 != nil {
		t.Errorf("Expected %v got %v", nil, nth10)
	}

	if nthMinus10 != nil {
		t.Errorf("Get first value expected %v got %v", nil, nthMinus10)
	}
}

func TestRemoveItem(t *testing.T) {
	l := createList()

	toDelItem := l.GetNth(1)
	toDelItem.Remove()

	if l.First().Next().Value() != i3 {
		t.Error("After remove middle item, first item does not point to last item")
	}
	if l.Last().Prev().Value() != i1 {
		t.Error("After remove middle item, last item does not point to first item")
	}

	l = createList()
	l.First().Remove()
	if l.First().Prev() != nil {
		t.Error("After remove first item, new first item prev should be nil")
	}
	if l.First().Value() != i2 {
		t.Errorf("After remove first item, new first item expected %v, got %v", i2, l.First().Value())
	}

	l = createList()
	l.Last().Remove()
	if l.Last().Next() != nil {
		t.Error("After remove last item, new last item next should be nil")
	}
	if l.Last().Value() != i2 {
		t.Errorf("After remove last item, new last item expected %v, got %v", i2, l.First().Value())
	}
}

func TestInsertAfter(t *testing.T) {
	l := createList()
	// ["test", 123, true]

	// we want insert i4="4th" after 123
	// new list ["test", 123, "4th", true]
	nth := l.InsertAfterNth(1, i4)
	if nth.Value() != i4 {
		t.Errorf("Wrong return: want %v, got %v", i4, nth.Value())
	}
	if nth.Prev().Value() != i2 {
		t.Errorf("Nth prev item: want %v, got %v", i2, nth.Prev().Value())
	}
	if nth.Next().Value() != i3 {
		t.Errorf("Nth next item: want %v, got %v", i3, nth.Next().Value())
	}

	prev := l.GetNth(1)
	if prev.Next().Value() != i4 {
		t.Errorf("Previous next item: want %v, got %v", i4, prev.Next().Value())
	}

	next := l.GetNth(3)
	if next.Prev().Value() != i4 {
		t.Errorf("Next previous item: want %v, got %v", i4, next.Prev().Value())
	}

	// insert one more at the end and one to the front
	nth = l.InsertAfterNth(10, i4)
	nthMinus := l.InsertAfterNth(-2, i4)
	// now we expect
	// ["4th", "test", 123, "4th", true, "4th"]
	if l.Len() != 6 {
		t.Errorf("List length after insertion: want %v, got %v", 6, l.Len())
	}

	if l.Last().Value() != i4 {
		t.Errorf("New last item: want %v, got %v", i4, l.Last().Value())
	}

	if nthMinus.Next().Value() != i1 {
		t.Errorf("New front item next: want %v, got %v", i1, nthMinus.Next().Value())
	}
}
