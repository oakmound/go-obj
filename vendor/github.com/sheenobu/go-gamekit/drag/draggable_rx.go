// generated by genrx -name RxDraggable -type Draggable drag.go; DO NOT EDIT

package drag

import (
	"sync"
	"time"
)

// RxDraggable is the reactive wrapper for Draggable
type RxDraggable struct {
	value Draggable
	lock  sync.RWMutex

	handles     chan int
	subscribers []chan<- Draggable
}

// NewRxDraggable creates a new reactive object for the initial value of Draggable
func NewRxDraggable(v Draggable) *RxDraggable {
	return &RxDraggable{
		value:   v,
		handles: make(chan int, 10),
	}
}

// Get gets the Draggable
func (rx *RxDraggable) Get() Draggable {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the Draggable and notifies subscribers
func (rx *RxDraggable) Set(v Draggable) {
	rx.lock.Lock()
	defer rx.lock.Unlock()
	rx.value = v

	for _, s := range rx.subscribers {
		if s != nil {
			s <- v
		}
	}
}

// Subscribe subscribes to changes on the Draggable
func (rx *RxDraggable) Subscribe() *RxDraggableSubscriber {

	c := make(chan Draggable)

	s := &RxDraggableSubscriber{
		C:      c,
		parent: rx,
	}

	rx.lock.Lock()
	select {
	case handle := <-rx.handles:
		s.handle = handle
		rx.subscribers[handle] = c
	default:
		rx.subscribers = append(rx.subscribers, c)
		s.handle = len(rx.subscribers) - 1
	}

	rx.lock.Unlock()

	return s
}

// RxDraggableSubscriber allows subscribing to the reactive Draggable
type RxDraggableSubscriber struct {
	C      <-chan Draggable
	handle int
	parent *RxDraggable
}

// Close closes the subscription
func (s *RxDraggableSubscriber) Close() {
	// remove from parent and close channel
	s.parent.lock.Lock()
	close(s.parent.subscribers[s.handle])
	s.parent.subscribers[s.handle] = nil
	s.parent.lock.Unlock()

	go func() {
		select {
		case s.parent.handles <- s.handle:
		case <-time.After(1 * time.Millisecond):
		}
	}()
}
