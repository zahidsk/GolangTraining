package main

import (
	"fmt"
	"time"
)

func getNumbers(c chan int)  {
	no := []int{10, 20, 30, 40, 50}
	for _, n := range no{
		fmt.Println("Sending no : ",n)
		c<-n
	}
	close(c)
}

func main(){

	bChan := make(chan string, 2)
	bChan <- "This Channel"
	bChan <- "is buffered"
	fmt.Println(<-bChan)
	bChan <- "3rd Packet"
	fmt.Println(<-bChan)
	fmt.Println(<-bChan)
	close(bChan)

	nChan := make(chan int, 2)
	go getNumbers(nChan)
	time.Sleep(2*time.Second)
	sum := 0
	for {
		num, ok :=<-nChan
		if ok==false{
			fmt.Println("Ohhhh  Channel is closed !!!")
			break
		}
		fmt.Println("Number received from GR is : ", num)
		sum = sum + num
		time.Sleep(2*time.Second)
	}
	fmt.Println("Sum of no. is : ",sum)
}
