package main

import "fmt"

func main() {
	slices := [][]int{{1, 2}, {2, 3}}
	result := SumAll(slices...)
	fmt.Println(result)
}

func SumFunction(i []int) int {
	var result int
	for _, num := range i {
		result += num
	}
	return result
}

func SumAll(numbers ...[]int) []int {
	newSlice := make([]int, len(numbers))
	for i, number := range numbers {
		for _, n := range number {
			newSlice[i] += n
		}
	}
	return newSlice
}
