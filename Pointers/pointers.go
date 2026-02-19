package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)
	fmt.Println("Adult years: ", getAdultYears(&age))
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	// overwrites the value at the pointer
	*age -= 18
	return *age
}
