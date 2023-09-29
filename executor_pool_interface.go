package executor

type ExecutorPool interface {
	Start()
	Submit(Task)
	// For for all submitted tasks to be executed
	Wait()
}
