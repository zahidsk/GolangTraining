package main

import (
	"fmt"
	"time"
)

func myprint()  {
	for i:=1; i<=5; i++{
		fmt.Println("Print from routine 1")
		time.Sleep(time.Duration(i)*time.Second)
	}
}

func main()  {
	fmt.Println("Print from main func")
	go myprint()

	for i:=5; i>=1; i--{
		fmt.Println("Print from main func")
		time.Sleep(time.Duration(i)*time.Second)
	}
}