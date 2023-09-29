package executor

type RRExecutorPool struct {
	taskQueue ConcurrentQueue[Task]
	executors []Executor
	cursor int
}

func NewRRExecutorPool(size int) *RRExecutorPool {
	es := make([]Executor, size)
	for i := range es {
		es[i] = NewSimpleExecutor()
	}
	return &RRExecutorPool{
		taskQueue: NewSimpleConcurrentQueue[Task](),
		executors: es,
		cursor: 0,
	}
}

func (rrep *RRExecutorPool) dispatch() {
	for {
		t := rrep.taskQueue.Dequeue()
		rrep.executors[rrep.cursor].Execute(t)
		rrep.cursor = (rrep.cursor + 1) % len(rrep.executors)
	}
}

func (rrep *RRExecutorPool) Start() {
	for _, e := range rrep.executors {
		e.Start()
	}
	go rrep.dispatch()
}

func (rrep *RRExecutorPool) Submit(t Task) {
	rrep.taskQueue.Enqueue(t)
}

func (rrep *RRExecutorPool) Wait() {
	rrep.taskQueue.Wait()
	for _, e := range rrep.executors {
		e.Wait()
	}
}
