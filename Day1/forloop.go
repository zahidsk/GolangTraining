package main

import "fmt"

func  main()  {
	StarForLoop()
	NumbForLoop()
}
// Exercise 2  : print star
func StarForLoop()  {
	for i:=1; i<=5; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// Exercise 3 : print number
func  NumbForLoop()  {
	count := 1
	for i:=1; i<=5; i++{
		for j:=1; j<=i; j++{
			fmt.Printf(" %d",count)
			count++
		}
		fmt.Println()
	}
}
