package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	// anonymous function that does the same thing as double, good for one-time uses
	fmt.Println(transformNumbers(&numbers, func(number int) int {
		return number * 2
	}))
	// passing the interface function "triple"
	fmt.Println(transformNumbers(&numbers, triple))
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, number := range *numbers {
		dNumbers = append(dNumbers, transform(number))
	}
	return dNumbers
}

//func double(number int) int {
//	return number * 2
//}

func triple(number int) int {
	return number * 3
}
