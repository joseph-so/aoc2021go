package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := -1
	previous := -9999

	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if current > previous {
			count += 1
		}
		previous = current
	}

	fmt.Println(count)
}
