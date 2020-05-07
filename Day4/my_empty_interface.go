package main

import "fmt"

type Mycity string

type Pincode int

type Emplpyee struct {
	fname string
	lname string
}

func Explain(i interface{})  {
	fmt.Printf("Recived value of type '%T' and having value '%v'\n",i,i)
}
func main()  {
	c := Mycity("Pune")
	pn := Pincode(23435)
	e := Emplpyee{"Raju", "Rastogi"}
	Explain(c)
	Explain(pn)
	Explain(e)
}