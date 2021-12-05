package main

import (
	"aoc2021go/src/readFile"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	data, err := readFile.ReadLine("input/day03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var stack [12][]int

	for _, record := range data {
		for i, bit := range record {
			stack[i] = append(stack[i], int(bit-'0'))
		}

	}

	var x string

	for i, _ := range stack {
		sort.Ints(stack[i])
		x += strconv.Itoa(stack[i][len(data)/2])
	}
	gamma, _ := strconv.ParseInt(x, 2, 16)
	epsilon := 0b111111111111 ^ gamma

	fmt.Printf("The power consumption of the submarine is %d\n", gamma*epsilon)
}
