package main

import "fmt"

type Employee struct {
	fname string
	lname string
	age int
	department string
}

func (e Employee)full_name() string {
	return e.fname+" "+e.lname
}
func (e *Employee)ChangeDepartment(newDept string) bool{
	e.department = newDept
	return true
}
func main()  {
	var e Employee=Employee{"Go", "Lang",10, "Programming"}
	fmt.Printf("%v\n",e)
	fn := e.full_name()
	fmt.Println("Full Name of emp : ",fn)
	e1 := &e
	e1.ChangeDepartment("CloudLanguage")
	fmt.Println("After changing the department ")
	fmt.Printf("%v\n",e)
}