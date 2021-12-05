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
	aim := 0

	for _, command := range data {
		parameters := strings.Fields(command)
		step, _ := strconv.Atoi(parameters[1])
		switch parameters[0] {
		case "forward":
			horizontal += step
			depth += step * aim
		case "down":
			aim += step
		case "up":
			aim -= step
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d, Result: %d\n", horizontal, depth, horizontal*depth)
}
