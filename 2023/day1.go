package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := io.ReadAll(file)
	arr := bytes.Split(raw, []byte("\n"))
	sum := 0
	var nums []int
	for _, line := range arr {
		var num_line []int
		for _, char := range line {
			if char >= 48 && char <= 57 {
				num_line = append(num_line, int(char)-48)
			}
		}
		if len(num_line) > 0 {
			num := num_line[0]*10 + num_line[len(num_line)-1]
			nums = append(nums, num)
			sum += num
		}
	}

	fmt.Println(nums)
	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := io.ReadAll(file)
	arr := bytes.Split(raw, []byte("\n"))
	sum := 0
	var nums []int
	for _, line := range arr {
		re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[1-9]`)
		matches := re.FindAll(line, -1)
		if len(matches) > 0 {
			number := convertToDigit(matches[0])*10 + convertToDigit(matches[len(matches)-1])
			nums = append(nums, number)
			sum += number
		}
	}
	numbersToMatch := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	type SearchResult struct {
		Index int
		Digit int
	}
	var res []SearchResult

	line := []byte(`twone`)
	for _, numToMatch := range numbersToMatch {
		re := regexp.MustCompile(numToMatch)
		matches := re.FindAll(line, -1)
		res = append(res, matches...)
	}
	// fmt.Println(sum)
	// fmt.Println(nums)
}

func convertToDigit(in []byte) int {
	if len(in) == 1 {
		return int(in[0]) - 48
	}

	switch string(in) {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}
