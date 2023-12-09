package main

import (
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) {
	result := 0
Loop:
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		sets := strings.Split(strings.Split(lines[i], ": ")[1], ";")
		for j := range sets {
			cubes := strings.Split(strings.Trim(sets[j], " "), ",")
			for k := range cubes {
				parsed := strings.Split(strings.Trim(cubes[k], " "), " ")
				color := parsed[1]
				count, _ := strconv.Atoi(parsed[0])
				if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
					continue Loop
				}
			}
		}
		result += i + 1
	}

	println(result)
}

func part2(lines []string) {
	result := 0
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		max := map[string]int{"red": 0, "green": 0, "blue": 0}
		sets := strings.Split(strings.Split(lines[i], ": ")[1], ";")
		for j := range sets {
			cubes := strings.Split(strings.Trim(sets[j], " "), ",")
			for k := range cubes {
				parsed := strings.Split(strings.Trim(cubes[k], " "), " ")
				color := parsed[1]
				count, _ := strconv.Atoi(parsed[0])
				if count > max[color] {
					max[color] = count
				}
			}
		}
		result += max["red"] * max["green"] * max["blue"]
	}

	println(result)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)

	part1(strings.Split(input, "\n"))
	part2(strings.Split(input, "\n"))
}
