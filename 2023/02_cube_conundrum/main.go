package main

import (
	// "fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"bufio"
)

type Color struct {
	name  string
	count int
}

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	num  int
	sets []Set
}

type Bag struct {
	red   int
	green int
	blue  int
}

func IsPossible(game *Game, bag *Bag) bool {
	for _, set := range game.sets {
		if set.red > bag.red || set.green > bag.green || set.blue > bag.blue {
			return false
		}
	}
	return true
}

func SumPossibleGames(r io.Reader, bag *Bag) int {
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		game := ParseGame(scanner.Text())
		if IsPossible(&game, bag) {
			sum += game.num
		}
	}
	return sum
}

func ParseColor(s string) Color {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	matches := re.FindStringSubmatch(s)
	// fmt.Printf("matches=%#v\n", matches)
	if len(matches) != 3 {
		panic("invalid match")
	}
	count, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return Color{
		name:  matches[2],
		count: count,
	}
}

func ParseSet(s string) Set {
	set := Set{}
	for _, c := range strings.Split(s, ",") {
		color := ParseColor(c)
		switch color.name {
		case "red":
			set.red = color.count
		case "green":
			set.green = color.count
		case "blue":
			set.blue = color.count
		default:
			panic("unknown color")
		}
	}
	return set
}

func ParseSets(s string) []Set {
	sets := []Set{}
	for _, set := range strings.Split(s, ";") {
		sets = append(sets, ParseSet(set))
	}
	return sets
}

func ParseGame(s string) Game {
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		panic("invalid match")
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return Game{
		num:  num,
		sets: ParseSets(matches[2]),
	}
}
