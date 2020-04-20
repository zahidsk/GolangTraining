package main

import (
	"fmt"
	"io/ioutil"
	//"os"
)

type contactInfo struct {
	address string
	phone   int
}
type student struct {
	fname string
	lname string
	class string
	age   int
	contactInfo
}

func myprint(st student) {
	fmt.Printf("\nFirstName: %s \t LastName: %s\t Class: %s\n", st.fname, st.lname, st.class)
	fmt.Printf("Address: %s \t Contact :%d", st.address, st.phone)
}
func updateStudentCont(st *student, cinfo contactInfo) {
	(*st).contactInfo = cinfo
}
func storeStudentRecord(st student, filename string) bool {
	str := fmt.Sprintf("FirstName: %s \t LastName: %s\t Class: %s\nAddress: %s \t Contact :%d", st.fname, st.lname, st.class, st.address, st.phone)
	//fmt.Printf(str)
	err := ioutil.WriteFile(filename, []byte(str), 0644)
	if err != nil {
		fmt.Println("Error while writing to file")
		return false
	}
	return true
}

func main() {
	st1 := student{"Shaikh", "Haris", "Engg", 20, contactInfo{"Pune", 9879874}}
	fmt.Printf("%v", st1)
	st1.contactInfo = contactInfo{"Mumbai", 9879874}
	myprint(st1)
	updateStudentCont(&st1, contactInfo{"Dhule", 124545})
	myprint(st1)
	storeStudentRecord(st1, "mystoredstruct.txt")
	//fmt.Println("Let's Start Struct")
}
