package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("This is first class function")
	}
	a()
	fmt.Printf("%T\n", a)
	fmt.Println("This is main function")
}

// Use Cases:
// 1. Can be used as decorators i.e passing function as a argument to other function
// 2. It can be use like lambda function in python
