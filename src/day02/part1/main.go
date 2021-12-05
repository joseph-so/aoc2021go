package main

import (
	"aoc2021go/src/readFile"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile.ReadLine("input/day02.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	horizontal := 0
	depth := 0

	for _, command := range data {
		parameters := strings.Fields(command)
		step, _ := strconv.Atoi(parameters[1])
		switch parameters[0] {
		case "forward":
			horizontal += step
		case "down":
			depth += step
		case "up":
			depth -= step
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d, Result: %d\n", horizontal, depth, horizontal*depth)
}
