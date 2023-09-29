package executor

import "sync"

type SimpleConcurrentQueue[T any] struct {
	items []T
	size sync.WaitGroup
	lock sync.Mutex
	waiting bool
	emptyLock sync.Mutex
}

func NewSimpleConcurrentQueue[T any]() *SimpleConcurrentQueue[T] {
	scq := &SimpleConcurrentQueue[T]{
		items: make([]T, 0),
	}
	scq.emptyLock.Lock()
	return scq
}

func (scq *SimpleConcurrentQueue[T]) Enqueue(item T) {
	if !scq.waiting {
		scq.lock.Lock()
		scq.items = append(scq.items, item)
		scq.lock.Unlock()
		if len(scq.items) == 1 {
			scq.emptyLock.Unlock()
		}
		scq.size.Add(1)
	}
}

func (scq *SimpleConcurrentQueue[T]) Dequeue() T {
	scq.emptyLock.Lock()
	scq.lock.Lock()
	item := scq.items[len(scq.items) - 1]
	scq.items = scq.items[:len(scq.items) - 1]
	scq.lock.Unlock()
	if len(scq.items) > 0 {
		scq.emptyLock.Unlock()
	}
	scq.size.Done()
	return item
}

func (scq *SimpleConcurrentQueue[T]) Wait() {
	scq.lock.Lock()
	scq.waiting = true
	scq.lock.Unlock()
	scq.size.Wait()
}
