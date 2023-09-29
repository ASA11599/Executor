package executor

type Executor interface {
	Start()
	Execute(Task)
	// Wait for the task to finish
	Wait()
}
