package examples

import (
	"GoExamples/async"
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func HelloWorldAtN(n int) {
	actions := make([]async.ActionWithData, n)
	for i := range actions {
		actions[i] = async.ActionWithData{
			Name: "helloWorld",
			Function: func() (interface{}, error) {
				helloWorld()
				return nil, nil
			},
		}
	}

	async.ParallelWithData(actions...)
}
