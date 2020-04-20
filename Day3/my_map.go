package main

import (
	"fmt"
	"math/rand"
)

type studentRec map[int]string

func createStud_record(st_name ...string) studentRec {
	st := studentRec{}
	for i := 1; i <= len(st_name); i++ {
		st[i] = st_name[i-1]
	}
	print_record(st)
	return st
}

func print_record(m map[int]string) {
	fmt.Println("Roll Number |\t Name")
	fmt.Println("-----------------------")
	for k, v := range m {
		fmt.Printf("\t%d  |\t%s\n", k, v)
	}
}

func (st studentRec) addUpdateStudentRec(rn int, name string) {
	st[rn] = name
}

func (st studentRec) removeFailedStud(rnos ...int) {
	for rn := 0; rn < len(rnos); rn++ {
		fmt.Printf("Deleting student with name '%s', and roll no. '%d'\n", st[rnos[rn]], rnos[rn])
		delete(st, rnos[rn])
	}
}

func main() {
	//m1 := map[int]string {}
	m1 := createStud_record("Raju", "Johan", "Jack", "Shaikh", "Tao", "chinu")
	fmt.Println(m1)
	fmt.Println("==================Adding======================")
	fmt.Println("Adding new Student Record")
	m1.addUpdateStudentRec(7, "Mack")
	print_record(m1)
	fmt.Println("==================Updating======================")
	fmt.Println("Updating name of roll no 1")
	m1.addUpdateStudentRec(1, "Kasar Jack")
	print_record(m1)
	fmt.Println("==================Deleting=====================")
	fmt.Println("Deleting the students who get failed...")
	fail_stud_roll := rand.Intn(7)
	fmt.Println("Failed student roll no.: ", fail_stud_roll)
	m1.removeFailedStud(fail_stud_roll)
	print_record(m1)
}
