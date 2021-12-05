package main

import (
	"aoc2021go/src/readFile"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile.ReadLine("input/day03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Strings(data)
	orgData := make([]string, len(data))
	crsData := make([]string, len(data))
	copy(orgData, data)
	copy(crsData, data)

	run := 0
	mask := ""
	for len(orgData) != 1 {
		filter := mask + "1" + strings.Repeat("0", 11-run)
		index := sort.Search(len(orgData), func(i int) bool { return orgData[i] >= filter })
		switch orgData[len(orgData)/2][run] {
		case '0':
			orgData = orgData[:index]
			mask += "0"
		case '1':
			orgData = orgData[index:]
			mask += "1"
		}
		run += 1
	}

	org, _ := strconv.ParseInt(orgData[0], 2, 16)

	run = 0
	mask = ""

	for len(crsData) != 1 {
		filter := mask + "1" + strings.Repeat("0", 11-run)
		index := sort.Search(len(crsData), func(i int) bool { return crsData[i] >= filter })
		switch crsData[len(crsData)/2][run] {
		case '0':
			crsData = crsData[index:]
			mask += "1"
		case '1':
			crsData = crsData[:index]
			mask += "0"
		}
		run += 1
	}

	crs, _ := strconv.ParseInt(crsData[0], 2, 16)
	fmt.Printf("Oxygen generator rating: %s (%d)\n", orgData[0], org)
	fmt.Printf("CO2 scrubber rating: %s (%d)\n", crsData[0], crs)
	fmt.Printf("Life support rating: %d", org*crs)

}
