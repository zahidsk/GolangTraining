package main

import "fmt"

type Car struct {
	color string
	noOfTyre int
	Company string
}

func (c Car)CarDetail()  {
	fmt.Printf("Car company is : %s Car coloris : %s" ,c.Company, c.color)
}


