package queue

import "fmt"

// Item is an interface to allow polymorphic objects
// to be placed into the queue. All they need to do is
// implement a String() identifier method and a Search()
// method which returns either more items to enqueue and
// search or a boolean declaration that it has found the
// Item we are looking for, and the found Item.
type Item interface {
	fmt.Stringer
	Search() ([]*Item, bool, *Item)
}

// Queue is a basic queue implementation without any
// synchronization. If 3 processes called ReadQueue they
// would have the same item and if they called dequeue
// two unread items would be lost.
//
// The queue is backed by a slice for order and a map
// whose keys are constructed via Item.String() for set
// membership tests.
type Queue struct {
	row    []*Item
	column map[string]bool
}

// NewQueue initializes the queue placing the parameter Item
// pointer as the first queue entry.
func NewQueue(itP *Item) *Queue {
	return &Queue{
		row:    []*Item{itP},
		column: map[string]bool{(*itP).String(): true},
	}
}

// enqueue will add item pointers to the queue as long as
// the item is not judged to already be present in the set
// of queue items.
func (q *Queue) enqueue(its []*Item) {
	for _, itP := range its {
		var itemId = (*itP).String()
		_, ok := q.column[itemId]
		if ok {
			continue
		}
		fmt.Println("Adding item", itemId, "to queue.")
		q.row = append(q.row, itP)
		q.column[itemId] = true
	}
}

// dequeue marks the current item at the head of the
// queue as done effectively removing it.
func (q *Queue) dequeue() {
	fmt.Println("Removing item", (*q.row[0]).String(), "from queue.")
	delete(q.column, (*q.row[0]).String())
	q.row = q.row[1:]
}

// read returns the next item pointer from
// the queue without changing the queue.
func (q *Queue) read() *Item {
	return q.row[0]
}

// Seek performs a queue-wide breadth-first search
// queue as done effectively removing it.
func (q *Queue) Seek() *Item {
	for len(q.row) > 0 {
		iPointers, found, foundPointer := (*q.read()).Search()
		if found {
			return foundPointer
		}
		q.enqueue(iPointers)
		q.dequeue()
	}
	panic("Queue unexpectedly has not found what it seeks.")
}
