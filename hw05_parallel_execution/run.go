package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {

	taskCh := make(chan Task)
	var wg sync.WaitGroup
	var errCnt int32

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for task := range taskCh {
				task()
			}
		}()
	}

	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	wg.Wait()
	return nil
}
