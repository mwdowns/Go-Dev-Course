package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, number := range *numbers {
		dNumbers = append(dNumbers, transform(number))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
