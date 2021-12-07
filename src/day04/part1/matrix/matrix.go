package matrix

import (
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type Matrix struct {
	values map[int]position
	Sum    int
	x      []int
	y      []int
}

func New(data []string) Matrix {
	var m Matrix
	m.x = []int{0, 0, 0, 0, 0}
	m.y = []int{0, 0, 0, 0, 0}
	m.Sum = 0
	m.values = make(map[int]position)
	re := regexp.MustCompile(`\s+`)
	for i, line := range data {
		values := re.Split(strings.TrimSpace(line), -1)
		for j, value := range values {
			v, _ := strconv.Atoi(value)
			m.Sum += v
			m.values[v] = position{x: j, y: i}
		}
	}

	return m
}

func (m *Matrix) Number(n int) bool {
	position, present := m.values[n]
	if present {
		m.Sum -= n
		m.x[position.x] += 1
		m.y[position.y] += 1
		if m.x[position.x] == 5 || m.y[position.y] == 5 {
			return true
		}
	}
	return false
}
