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

const debug = false

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
	re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	scanner := bufio.NewScanner(r)
	var i, sum = 0, 0
	for scanner.Scan() {
		i++
		// Replace "one" with "o1e", "three" with "t3e" and so on
		replaced := re.ReplaceAllStringFunc(scanner.Text(), getReplacement)
		matches := re.FindAllString(replaced, -1)
		if len(matches) == 0 {
			return 0, errors.New("no matches found")
		}

		first, last := matches[0], matches[len(matches)-1]

		value, err := strconv.Atoi(getStringValue(first) + getStringValue(last))
		if err != nil {
			return 0, err
		}
		sum += value
		if debug {
			fmt.Printf("[%d] line=%s, matches=%#v, value=%d, sum=%d\n", i, scanner.Text(), matches, value, sum)
		}
	}

	return sum, nil
}

func getReplacement(s string) string {
	replacements := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	if replacement, ok := replacements[s]; ok {
		return replacement
	}
	return s
}

func getStringValue(s string) string {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if num, ok := nums[s]; ok {
		return strconv.Itoa(num)
	}
	return s
}
