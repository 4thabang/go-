package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	slc := []int{1, 2, 3, 4, 5}
	result := SumMapReduce(slc)
	fmt.Println(result)

	fileMap := map[string]string{
		"test.txt":   "Hello, world",
		"noTest.txt": "Bye, world",
	}

	for fileName, fileData := range fileMap {
		if err := CreateFile(fileName, fileData); err != nil {
			panic(err)
		}
	}
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

func CreateFile(fileName string, data string) error {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := WriteToFile(file, data); err != nil {
		return err
	}
	return nil
}

func WriteToFile(w io.Writer, data string) error {
	_, err := w.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
