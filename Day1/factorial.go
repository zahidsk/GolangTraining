package main

import "fmt"

func main()  {
	n := 5
	//ret = factorial(n)
	fmt.Printf("Factorial of a number %d is %d ", n, factorial(n))
}

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return factorial(n-1)*n
}