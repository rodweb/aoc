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

func TestFirstPart(t *testing.T) {
	t.Run("Works with first part", func(t *testing.T) {
		test(t, "input.txt", 54304)
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
