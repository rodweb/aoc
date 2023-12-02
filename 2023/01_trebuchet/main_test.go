package main

import (
	"bufio"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("Works with provided example", func(t *testing.T) {
		file, err := os.Open("example.txt")
		if err != nil {
			t.Fatal(err)
		}
	        reader := bufio.NewReader(file)

		got, err := SumCalibrationValues(reader)
		if err != nil {
			t.Fatal(err)
		}

		want := 142
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
