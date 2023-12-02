package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestFirstExample(t *testing.T) {
	t.Run("Works with first example", func(t *testing.T) {
		test(t, fromFile("first_example.txt"), 142)
	})
}

func TestSecondExample(t *testing.T) {
	t.Run("Works with second example", func(t *testing.T) {
		test(t, fromFile("second_example.txt"), 281)
	})
}

func TestSpelledLetters(t *testing.T) {
	t.Run("Works with spelled letters", func(t *testing.T) {
		test(t, fromString("one\n"), 11)
		test(t, fromString("two\n"), 22)
		test(t, fromString("three\n"), 33)
		test(t, fromString("four\n"), 44)
		test(t, fromString("five\n"), 55)
		test(t, fromString("six\n"), 66)
		test(t, fromString("seven\n"), 77)
		test(t, fromString("eight\n"), 88)
		test(t, fromString("nine\n"), 99)
	})
}

func TestOverlapped(t *testing.T) {
	t.Run("Works with overlap", func(t *testing.T) {
		test(t, fromString("oneeighthree\n"), 13)
	})
}

func TestInput(t *testing.T) {
	t.Run("Works with input file", func(t *testing.T) {
		test(t, fromFile("input.txt"), 54418)
	})
}

func fromFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func fromString(s string) io.Reader {
	return strings.NewReader(s)
}

func test(t *testing.T, reader io.Reader, want int) {
	got, err := SumCalibrationValues(reader)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
