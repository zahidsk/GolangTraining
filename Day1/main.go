package main

import "fmt"

func main() {
	fmt.Println("Hello Go with GoLand")

	m := map[string]int{"one": 1, "two": 2}
	for x, y := range m {
		//continue
		fmt.Println("Hello")

		fmt.Printf("%s, %d", x, y)
	}
}
