package practice

import (
	"errors"
	"fmt"
	"sync"
)

type Action func() error

var n = 0

func Actionable(str string, n int) *int {
	fmt.Println(str)
	n = n + 1
	return &n
}

func ActionableAtN(n int) error {
	actions := make([]Action, n)
	for i := range actions {
		actions[i] = func() error {
			n = *Actionable("Hello World", n)
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
	return n
}

func Parallel(actions ...Action) error {
	var wg sync.WaitGroup
	wg.Add(len(actions))

	for _, action := range actions {
		go func(a Action) { // Create a copy of action
			defer func() {
				if r := recover(); r != nil {
					_ = errors.New("panic recovered")
				}
				wg.Done() // Ensure wg.Done() is always called
			}()
			err := a()
			if err != nil {
				return
			} // Actually execute the action
		}(action)
	}

	wg.Wait()
	return nil
}
