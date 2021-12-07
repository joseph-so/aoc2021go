package main

import (
	"aoc2021go/src/day04/part1/matrix"
	"aoc2021go/src/readFile"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile.ReadLine("input/day04.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	drawNumbers := data[0]
	var boards []matrix.Matrix

	for i := 2; i < len(data)-1; i += 6 {
		boards = append(boards, matrix.New(data[i:i+5]))
	}

	for _, n := range strings.Split(drawNumbers, ",") {
		number, _ := strconv.Atoi(n)
		for i := range boards {
			if boards[i].Number(number) {
				fmt.Println("number: ", number)
				fmt.Println("sum: ", boards[i].Sum)
				fmt.Println("The final score is ", boards[i].Sum*number)
				return
			}
		}
	}
}
