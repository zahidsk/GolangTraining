package main

import "fmt"

func myCities() []string {
	s1 := []string{"Pune", "Mumbai", "Dhule", "Kolkata", "Banglore", "Mysore", "Pilkod"}
	return s1
}
func UpdateCityatIndex(s []string, ncity string, index int) {
	//s = append(s, ncity)
	s[index] = ncity
}

func main() {
	s1 := myCities()
	//in1 := []int{1,2,3}
	//fmt.Println("My Int slice",in1)

	fmt.Println("My String slice", s1)
	s1 = append(s1, "Delhi")

	// length of slice
	fmt.Printf("Total Number of cities in slice: %d\n", len(s1))

	// update at index
	UpdateCityatIndex(s1, "Chennai", 1)
	fmt.Println("After update my String slice\n", s1)
	fmt.Printf("%s", s1[3])

	// Copy and update new slice
	s2 := make([]string, len(s1))
	copy(s2, s1)
	fmt.Printf("\n S2 City list : %s\n", s2)
	UpdateCityatIndex(s2, "Kondwa", 5)
	fmt.Printf("After update S2 City list : %s\n", s2)
	fmt.Printf("After update S1 City list : %s\n", s1)
}
