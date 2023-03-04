package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Running tests: [START]")
	exit := m.Run()
	log.Println("Ending tests: [END]")

	os.Exit(exit)
}

func TestSumFunction(t *testing.T) {
	tables := []struct {
		name  string
		want  int
		got   int
		input []int
	}{
		{name: "low slice", input: []int{1, 2, 3, 4, 5}},
		{name: "high slice", input: []int{333, 232, 98}},
		{name: "negative slice", input: []int{-10, -5, -50, -4}},
		{name: "empty slice", input: []int{}},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			for _, num := range tt.input {
				tt.want += num
			}
			tt.got = SumFunction(tt.input)
			assertHandler(t, tt.input, tt.got, tt.want)
		})
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, expected: %d", got, want)
	}
}

func assertHandler(t testing.TB, input []int, got, want int) {
	t.Helper()
	if want != got {
		t.Errorf("expect: %d, got: %d, input: %v", want, got, input)
	}
}
