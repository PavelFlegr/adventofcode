package main

import (
	"os"
	"strconv"
	"strings"
)

func getLineValue(line []int) (int, int) {
	done := true
	for _, val := range line {
		if val != 0 {
			done = false
			break
		}
	}
	if done {
		return 0, 0
	}

	var next []int
	for i := range line {
		if i == 0 {
			continue
		}
		next = append(next, line[i]-line[i-1])
	}

	first, last := getLineValue(next)
	return line[0] - first, last + line[len(line)-1]
}

func part1and2(lines []string) {
	prevSum := 0
	nextSum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		var vals []int
		for _, val := range strings.Split(line, " ") {
			parsed, _ := strconv.Atoi(val)
			vals = append(vals, parsed)
		}
		prev, next := getLineValue(vals)
		prevSum += prev
		nextSum += next
	}

	println(nextSum, prevSum)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")

	part1and2(lines)
}
