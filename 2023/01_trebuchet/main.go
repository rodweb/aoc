package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	sum, err := SumCalibrationValues(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum=%d\n", sum)
}

func SumCalibrationValues(r io.Reader) (int, error) {
	var values []int
	re := regexp.MustCompile(`(\d)`)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		switch len(matches) {
		case 0:
			return 0, errors.New("no matches found")
		case 1:
			value, err := strconv.Atoi(matches[0] + matches[0])
			if err != nil {
				return 0, err
			}
			values = append(values, value)
		default:
			value, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
			if err != nil {
				return 0, err
			}
			values = append(values, value)
		}
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, nil
}
