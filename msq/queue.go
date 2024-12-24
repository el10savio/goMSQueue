package queue

import (
	"sync/atomic"
)

type Queue[T any] struct {
	head atomic.Pointer[Node[T]]
	tail atomic.Pointer[Node[T]]
}

type Node[T any] struct {
	element *T
	next    atomic.Pointer[Node[T]]
}

func NewMSQueue[T any]() *Queue[T] {
	queue := &Queue[T]{}

	dummy := newNode[T](nil)
	queue.head.Store(dummy)
	queue.tail.Store(dummy)

	return queue
}

func newNode[T any](element *T) *Node[T] {
	return &Node[T]{element: element}
}

// Add to the tail
func (q *Queue[T]) Enqueue(element T) {
	newElement := newNode[T](&element)

	for {
		// get the tail and its next node
		tail := q.tail.Load()
		next := tail.next.Load()

		// check if our tail is the latest
		// otherwise retry
		if tail == q.tail.Load() {
			// if tail's next is not empty means
			// we're still lagging so advance
			if next == nil {
				// again check if tail's next is empty
				// and then fill it in
				if tail.next.CompareAndSwap(nil, newElement) {
					// and now the tail becomes
					// the newly inserted element
					q.tail.CompareAndSwap(tail, newElement)
					return
				}
			} else {
				q.tail.CompareAndSwap(tail, next)
			}
		}
	}
}

// Remove from the head
func (q *Queue[T]) Dequeue() *T {
	for {
		// get the tail, head, and its next node
		tail := q.tail.Load()
		head := q.head.Load()
		next := head.next.Load()

		// check if our head is the latest
		// otherwise retry
		if head == q.head.Load() {
			// check if the queue is empty
			if head == tail {
				// Confirm that the next
				// for an empty queue is nil
				if next == nil {
					return nil
				}
				// Otherwise, tail is lagging, so advance it
				q.tail.CompareAndSwap(tail, next)
			} else {
				// read value
				value := next.element
				// advance head
				if q.head.CompareAndSwap(head, next) {
					return value
				}
			}
		}
	}
}
