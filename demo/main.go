package main

import (
	"fmt"
	"io"
)

func main() {
	slc := []int{1, 2, 3, 4, 5}
	result := SumMapReduce(slc)
	fmt.Println(result)
}

func SumMapReduce(nums []int) int {
	result := reduce(nums, 0, iterator)
	return result
}

func reduce(arrs []int, initValue int, fn func(first, second int) int) int {
	for _, arr := range arrs {
		initValue = fn(initValue, arr)
	}
	return initValue
}

func iterator(initValue, value int) int {
	initValue += value
	return initValue
}

func WriteToFile(w io.Writer) error {
	return nil
}
