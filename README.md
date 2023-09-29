# Executor

A thread pool library

## Usage

```go

// Initialize a thread pool with 100 workers
var ep executor.ExecutorPool = executor.NewRRExecutorPool(100)

// Start the workers
ep.Start()

for i := 0; i < 10; i++ {
    n := (i + 1) * (i + 1)
    // Submit a task
    ep.Submit(func() {
        time.Sleep(2 * time.Second)
        fmt.Println(n)
    })
}

// Wait for all the tasks to finish
ep.Wait()
```
