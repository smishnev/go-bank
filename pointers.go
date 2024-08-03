package main

import "fmt"

func pointer() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)

	fmt.Println("Adult years: ", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
	// return *age - 18
}
