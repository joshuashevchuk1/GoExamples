package practice

import (
	"fmt"
	"sync"
)

type Action func() error

var (
	n  = 0
	mu sync.Mutex // Mutex for safe concurrent updates
)

func Actionable(str string) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(str)
	n++ // Modify the global variable safely
}

func ActionableAtN(count int) error {
	actions := make([]Action, count)
	for i := range actions {
		actions[i] = func() error {
			Actionable("Hello World") // Use global `n` safely
			return nil
		}
	}

	err := Parallel(actions...)
	if err != nil {
		return err
	}
	return nil
}

func GetN() int {
	mu.Lock()
	defer mu.Unlock()
	return n
}

func Parallel(actions ...Action) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(actions))

	for _, action := range actions {
		wg.Add(1)
		go func(a Action) {
			defer wg.Done()
			if err := a(); err != nil {
				errChan <- err
			}
		}(action)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err // Return the first error encountered
	}
	return nil
}
