package executor

import "sync"

type SimpleExecutor struct {
	taskQueue ConcurrentQueue[Task]
	taskWg sync.WaitGroup
}

func NewSimpleExecutor() *SimpleExecutor {
	return &SimpleExecutor{
		taskQueue: NewSimpleConcurrentQueue[Task](),
	}
}

func (se *SimpleExecutor) listen() {
	for {
		task := se.taskQueue.Dequeue()
		se.taskWg.Add(1)
		task()
		se.taskWg.Done()
	}
}

func (se *SimpleExecutor) Start() {
	go se.listen()
}

func (se *SimpleExecutor) Execute(t Task) {
	se.taskQueue.Enqueue(t)
}

func (se *SimpleExecutor) Wait() {
	se.taskQueue.Wait()
	se.taskWg.Wait()
}
