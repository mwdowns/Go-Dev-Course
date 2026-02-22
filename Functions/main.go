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

	// using closures
	fmt.Println(transformNumbers(&numbers, createTransformer(4)))

	// recursion in go
	fmt.Println(factorial(5))
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
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

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
	//result := 1
	//for i := 1; i < number; i++ {
	//	result = result * i
	//}
	//return result
}
