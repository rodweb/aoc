package main

import (
	"bufio"
	"os"
	"testing"
)

func TestFirstExample(t *testing.T) {
	t.Run("Works with first example", func(t *testing.T) {
		test(t, "first_example.txt", 142)
	})
}

func TestSecondExample(t *testing.T) {
	t.Run("Works with second example", func(t *testing.T) {
		test(t, "second_example.txt", 281)
	})
}

func TestInput(t *testing.T) {
	t.Run("Works with input file", func(t *testing.T) {
		test(t, "input.txt", 54412)
	})
}

func test(t *testing.T, filename string, want int) {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	reader := bufio.NewReader(file)

	got, err := SumCalibrationValues(reader)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
