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
	for _, game := range GetGames(r) {
		if IsPossible(&game, bag) {
			sum += game.num
		}
	}
	return sum
}

func SumPowerMinSet(r io.Reader) int {
	sum := 0
	for _, game := range GetGames(r) {
		minSet := MinSet(&game)
		power := minSet.red * minSet.green * minSet.blue
		sum += power
	}
	return sum
}

func MinSet(game *Game) Set {
	minSet := game.sets[0]
	for _, set := range game.sets {
		if set.red > minSet.red {
			minSet.red = set.red
		}
		if set.green > minSet.green {
			minSet.green = set.green
		}
		if set.blue > minSet.blue {
			minSet.blue = set.blue
		}
	}
	return minSet
}

func GetGames(r io.Reader) []Game {
	games := []Game{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		game := ParseGame(scanner.Text())
		games = append(games, game)
	}
	return games
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
