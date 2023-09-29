package main

import (
	"fmt"
	"time"

	"github.com/ASA11599/executor"
)

func main() {
	var ep executor.ExecutorPool = executor.NewRRExecutorPool(100)
	ep.Start()
	defer ep.Wait()
	for i := 0; i < 10; i++ {
		n := (i + 1) * (i + 1)
		ep.Submit(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(n)
		})
	}
}
