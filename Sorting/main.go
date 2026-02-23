// codility test that I was given
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	a1 := []int{1, 2, 3}                                   // 4
	a2 := []int{-10, -2000}                                // 1
	a3 := []int{7, 10, -1, 4, 1, 3, 1, 2}                  // 5
	a4 := []int{10, -3, 2, 7, 3}                           // 1
	a5 := []int{-7, 99, 2, 2, 5, 3, 1, 4, 7, -4, 11, 6, 5} // 8

	ars := [][]int{a1, a2, a3, a4, a5}
	for _, ar := range ars {
		fmt.Println(findNum(ar))
	}
}

func findNum(a []int) int {
	var result int
	var tempResult int
	slices.SortFunc(a, func(a int, b int) int {
		return cmp.Compare(a, b)
	})
	fmt.Println(a)
	for _, num := range a {
		if num <= 0 {
			result = 1
			continue
		} else if num > 1 && !(slices.Contains(a, 1)) {
			result = 1
			break
		} else {
			tempResult = num + 1
			result = checkInclusion(a, tempResult)
			break
		}
	}
	return result
}
func checkInclusion(a []int, tempResult int) int {
	if slices.Contains(a, tempResult) {
		return checkInclusion(a, tempResult+1)
	}
	return tempResult
}
