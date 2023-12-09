package main

import (
	"math"
	"os"
	"strings"
)

func contains(items []string, item string) bool {
	for i := range items {
		if items[i] == item {
			return true
		}
	}

	return false
}

func part1(lines []string) {
	sum := 0
	for r := range lines {
		p := 0
		if lines[r] == "" {
			continue
		}

		data := strings.Split(lines[r], ": ")[1]
		groups := strings.Split(data, " | ")
		winning := strings.Split(groups[0], " ")
		owned := strings.Split(groups[1], " ")

		for i := range winning {
			if winning[i] != "" && contains(owned, winning[i]) {
				p += 1
			}
		}
		sum += int(math.Pow(2, float64(p)-1))
	}
	println(int(sum))
}

func part2(lines []string) {
	copies := map[int]int{}
	for r := range lines {
		p := 0
		if lines[r] == "" {
			continue
		}

		data := strings.Split(lines[r], ": ")[1]
		groups := strings.Split(data, " | ")
		winning := strings.Split(groups[0], " ")
		owned := strings.Split(groups[1], " ")

		for i := range winning {
			if winning[i] != "" && contains(owned, winning[i]) {
				p += 1
			}
		}
		copies[r] += 1
		for i := 1; i <= p; i++ {
			copies[r+i] += 1 * copies[r]
		}
	}
	sum := 0
	for i := range copies {
		sum += copies[i]
	}
	println(int(sum))
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)

	part1(strings.Split(input, "\n"))
	part2(strings.Split(input, "\n"))
}
