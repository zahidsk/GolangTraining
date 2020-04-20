package main

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {
	st1 := student{"Shaikh", "Haris", "Engg", 18, contactInfo{"Pune", 9879874}}
	if (st1.age < 18) && (st1.class == "Engg") {
		t.Error("Engg student should be matured, age should be > 18")
	}
}

func TestContactInfo(t *testing.T) {
	st1 := student{"Shaikh", "Haris", "Engg", 15, contactInfo{"Pune", 989874129}}
	l := len(string(fmt.Sprintf("%d", st1.contactInfo.phone)))
	fmt.Println(l)
	if l != 10 {
		t.Error("Invalid Phone number. Phone number should be of 10 digit")
	}

}
