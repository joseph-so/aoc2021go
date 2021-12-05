package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLine(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		output = append(output, value)
	}

	return output, scanner.Err()
}

func main() {
	data, _ := readLine("input/day01.txt")
	count := 0
	for i := 3; i < len(data); i++ {
		if data[i] > data[i-3] {
			count += 1
		}
	}
	fmt.Println(count)
}
