package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type PatternLine []byte
type Pattern []PatternLine

func main() {
	part1()
}

func part1() {
	file, err := os.Open("input/day13_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := io.ReadAll(file)
	arr := bytes.Split(raw, []byte("\n"))
	patterns := make([]Pattern, 1)
	for _, v := range arr {
		if len(v) == 1 && int(v[0]) == 13 {
			patterns = append(patterns, make(Pattern, 0))
			continue
		}
		if len(v) > 1 && v[len(v)-1] == byte(13) {
			v = v[:len(v)-1]
		}
		patterns[len(patterns)-1] = append(patterns[len(patterns)-1], v)
	}
	// fmt.Println(patterns[0])
	mirrorPattern(patterns[0])
}

func transpose(input Pattern) Pattern {
	newArr := make(Pattern, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			newArr[j] = append(newArr[j], input[i][j])
		}
	}
	return newArr
}

func mirrorPattern(in Pattern) Pattern {
	out := make(Pattern, 0)
	for _, p := range in {
		reverseLine := make(PatternLine, len(p))
		copy(reverseLine, p)
		slices.Reverse(reverseLine)
		out = append(out, reverseLine)
	}
	return out
}

func autoconvolve(a Pattern) int {
	width := len(a[0])
	for i := 2; i < width; i++ {
		fmt.Println(i)
	}
	return -1
}
