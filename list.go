package list

// List is a container of double linked items
type List struct {
	//data  map[*Item]bool
	length int
	first  *Item
	last   *Item
}

//NewList constructor
func NewList() *List {
	return &List{}
}

// Len list length
func (l List) Len() int {
	return l.length
}

// First item in list
func (l List) First() *Item {
	return l.first
}

// Last item from list
func (l List) Last() *Item {
	return l.last
}

// PushFront item at the beginning of list
func (l *List) PushFront(v interface{}) *Item {
	i := &Item{value: v, container: l}

	if l.first == nil && l.last == nil {
		l.first = i
		l.last = i
		l.length++
		return i
	}

	first := l.first

	first.prev = i
	i.next = first
	l.first = i
	l.length++
	return i
}

// PushBack item to the end of list
func (l *List) PushBack(v interface{}) *Item {
	i := &Item{value: v, container: l}

	if l.first == nil && l.last == nil {
		l.first = i
		l.last = i
		l.length++
		return i
	}

	last := l.last

	last.next = i
	i.prev = last
	l.last = i
	l.length++
	return i
}

// GetNth return item at nth-index beginning from 0;
// if n < 0 we start iteration from last element;
// if n > list length we return nil
func (l List) GetNth(n int) *Item {
	if n >= 0 { // we iterate from first onwards
		if n >= l.length {
			return nil
		}
		i := l.first

		for j := 0; j < n; j++ {
			i = i.next
		}
		return i
	}
	// else we iterate backwads from last item
	if n <= -l.length {
		return nil
	}
	i := l.last
	for j := 0; j > n; j-- {
		i = i.prev
	}
	return i
}

// InsertAfterNth insert new item after nth element;
// if n >= length of list, then add back
// if n is negative add front
func (l *List) InsertAfterNth(n int, v interface{}) *Item {
	if n >= l.length {
		return l.PushBack(v)
	}
	if n < 0 {
		return l.PushFront(v)
	}

	nth := l.GetNth(n)
	next := nth.Next()
	i := &Item{value: v, container: l, prev: nth, next: next}
	nth.next = i
	next.prev = i

	l.length++

	return i
}

// Item is one item in double linked list
type Item struct {
	value     interface{}
	next      *Item
	prev      *Item
	container *List
}

// Value return value of current item
func (i Item) Value() interface{} {
	return i.value
}

// Next returns pointer to next Item
func (i Item) Next() *Item {
	return i.next
}

// Prev returns pointer to previous Item
func (i Item) Prev() *Item {
	return i.prev
}

// Remove self destruct
func (i *Item) Remove() {
	prev := i.Prev()
	next := i.Next()
	if prev == nil { // we remove first item
		next.prev = nil
		i.container.first = next
	} else if next == nil { // we remove last item
		prev.next = nil
		i.container.last = prev
	} else { // we remove middle item
		// we cross reference neighbor Items
		next.prev = prev
		prev.next = next
	}
	i.value = nil
	i.next = nil
	i.prev = nil
	i.container.length--
	i = nil
}
