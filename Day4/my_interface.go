package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	rd float32
}

func (ca Circle)Area() float32 {
	ar := math.Pi*(ca.rd*ca.rd)
	fmt.Println("area of Circle for radius ", ca.rd, " is : ", ar)
	return ar
}
func (ca Circle)Perimeter() float32 {
	pr := 2*math.Pi*ca.rd
	fmt.Println("Perimeter of Circle is : ", pr)
	return pr
}

func main()  {
	var s Shape
	s = Circle{5}
	s.Area()
	c := Circle{3}
	c.Area()
	c.Perimeter()
	fmt.Printf("type of s : %T\n",s)
	fmt.Printf("value of s : %v\n",s)
	fmt.Printf("type of C : %T\n",c)
	fmt.Printf("value of C : %v\n",c)
}
