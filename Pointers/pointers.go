package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age) // should be 32
	fmt.Println("Adult years: ", getAdultYears(&age))
	fmt.Println(age) // should be 14
}

func getAdultYears(age *int) int {
	// overwrites the value at the pointer
	*age -= 18
	return *age
}
