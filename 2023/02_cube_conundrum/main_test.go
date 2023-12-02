package main

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseColor(t *testing.T) {
	t.Run("Parse color", func(t *testing.T) {
		got := ParseColor("3 blue")
		want := Color{name: "blue", count: 3}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestParseSet(t *testing.T) {
	t.Run("Parse set", func(t *testing.T) {
		got := ParseSet("3 blue, 4 red")
		want := Set{red: 4, blue: 3}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestParseGame(t *testing.T) {
	t.Run("Parses game number", func(t *testing.T) {
		got := ParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
		want := Game{num: 1, sets: []Set{
			{red: 4, blue: 3},
			{red: 1, green: 2, blue: 6},
			{green: 2},
		}}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestIsPossible(t *testing.T) {
	bag := &Bag{red: 1, green: 2, blue: 3}
	t.Run("Game is possible", func(t *testing.T) {
		game := &Game{sets: []Set{
			{red: 1, blue: 3},
		}}
		got := IsPossible(game, bag)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("Game is not possible", func(t *testing.T) {
		game := &Game{sets: []Set{
			{red: 1, blue: 3},
			{red: 3, blue: 1},
		}}
		got := IsPossible(game, bag)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestFirstPart(t *testing.T) {
	bag := &Bag{red: 12, green: 13, blue: 14}
	t.Run("Example works for first part", func(t *testing.T) {
		got := SumPossibleGames(fromFile("example_01.txt"), bag)
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSecondPart(t *testing.T) {
	t.Run("Example works for second part", func(t *testing.T) {
		got := SumPowerMinSet(fromFile("example_01.txt"))
		want := 2286

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestInputFirstPart(t *testing.T) {
	bag := &Bag{red: 12, green: 13, blue: 14}
	t.Run("Input works for first part", func(t *testing.T) {
		got := SumPossibleGames(fromFile("input.txt"), bag)
		want := 2447

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestInputSecondPart(t *testing.T) {
	t.Run("Input works for second part", func(t *testing.T) {
		got := SumPowerMinSet(fromFile("input.txt"))
		want := 56322

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
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

func test(t *testing.T, reader io.Reader, bag *Bag, want int) {
}
