package main

import (
	"GoExamples/practice"
)

func main() {
	//examples.HelloWorldAtN(5)
	err := practice.ActionableAtN(5)
	if err != nil {
		return
	}
	n := practice.GetN()
	println(n)
}
