package executor

type ConcurrentQueue[T any] interface {
	Enqueue(T)
	Dequeue() T
	// Wait for all the items to be dequeued
	Wait()
}
