package main

import "fmt"
// Unbuffered channel

func sum(x []int, fc chan int)  {
	fmt.Println(x)
	summ := 0
	for _, n :=range x{
		//fmt.Println("n: ",n)
		summ +=n
	}
	fmt.Println("Result is ", summ)
	fc<-summ
}

func main()  {
	num := []int{1,3,5,7,9,2,4,6,8,10}
	fmt.Println("Length of a list : ",len(num))
	c := make(chan int)
	go sum(num[:len(num)/2], c)
	s1 := <-c
	fmt.Println("Sum1 : ", s1)
	go sum(num[len(num)/2:], c)
	s2 := <-c
	fmt.Println("Sum2 : ", s2)
	fmt.Println("Result of all numb from list : ", s1+s2)
	close(c)
}
