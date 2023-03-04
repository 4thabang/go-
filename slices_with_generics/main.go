package main

import "fmt"

type Typer interface {
	string | int
}

func main() {
	value := SumFunction([]int{1, 2, 3, 4})
	fmt.Println(value)

	slices := [][]int{{1, 2}, {2, 3}}
	result := SumAll(slices...)
	fmt.Println(result)
}

func SumFunction(i []int) int {
	sum := reduce(i, 0, reducer[int])
	return sum
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

func reduce[T Typer](arr []T, initValue T, fn func(first, second T) T) T {
	for _, value := range arr {
		initValue = fn(initValue, value)
	}
	return initValue
}

func reducer[T Typer](initValue T, value T) T {
	initValue = initValue + value
	return initValue
}
