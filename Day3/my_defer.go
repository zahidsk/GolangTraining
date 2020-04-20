package main

import "fmt"

func main() {
	defer fmt.Println("The End")
	defer fmt.Println("Please check locks")

	fmt.Println("Hi Good Morning ")
	fmt.Println("Start Car ..let's go for long drive")
}
