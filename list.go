package list

// List is a container of double linked items
type List struct {
	data  map[*Item]bool
	first *Item
	last  *Item
}

//NewList constructor
func NewList() *List {
	return &List{data: make(map[*Item]bool)}
}

// Len list length
func (l List) Len() int {
	return len(l.data)
}

// First item in list
func (l List) First() *Item {
	if l.Len() > 0 {
		return l.first
	}
	return nil
}

// Last item from list
func (l List) Last() *Item {
	if l.Len() > 0 {
		return l.last
	}
	return nil
}

// PushFront item at the beginning of list
func (l *List) PushFront(v interface{}) *Item {
	i := Item{value: v, container: l}

	if l.insertFirst(&i) {
		return &i
	}

	first := l.First()

	first.prev = &i
	i.next = first
	l.append(&i)
	l.first = &i
	return &i
}

// PushBack item to the end of list
func (l *List) PushBack(v interface{}) *Item {
	i := Item{value: v, container: l}

	if l.insertFirst(&i) {
		return &i
	}

	last := l.Last()

	last.next = &i
	i.prev = last
	l.append(&i)
	l.last = &i
	return &i
}

// GetNth return item at nth-index beginning from 0;
// if n < 0 we start iteration from last element;
// if n > list length we return last element
func (l List) GetNth(n int) *Item {
	if n >= 0 { // we iterate from first onwards
		if n >= l.Len() {
			return l.Last()
		}
		i := l.First()

		for j := 0; j < n; j++ {
			i = i.next
		}
		return i
	}
	// else we iterate backwads from last item
	if n <= -l.Len() {
		return l.First()
	}
	i := l.Last()
	for j := 0; j > n; j-- {
		i = i.prev
	}
	return i
}

// InsertAfterNth insert new item after nth element;
// if n >= length of list, then add back
// if n is negative add front
func (l *List) InsertAfterNth(n int, v interface{}) *Item {
	if n >= l.Len() {
		return l.PushBack(v)
	}
	if n < 0 {
		return l.PushFront(v)
	}

	nth := l.GetNth(n)
	next := nth.Next()
	i := Item{value: v, container: l, prev: nth, next: next}
	nth.next = &i
	next.prev = &i

	l.append(&i)

	return &i
}

func (l *List) append(i *Item) {
	l.data[i] = true
}

func (l *List) insertFirst(i *Item) bool {
	if (l.First() == nil) && (l.Last() == nil) { // this is empty list
		l.first = i
		l.last = i
		l.append(i)
		return true
	}
	return false
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
	delete(i.container.data, i)
	i = nil
}
