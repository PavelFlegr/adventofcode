package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	var err error
	for _, line := range lines {
		first := 0
		for i := range line {
			first, err = strconv.Atoi(string(line[i]))
			if err == nil {
				break
			}
		}

		last := 0
		for i := range line {
			last, err = strconv.Atoi(string(line[len(line)-i-1]))
			if err == nil {
				break
			}
		}

		sum += first*10 + last
	}

	println(sum)
}

func part2(input string) {
	candidates := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lines := strings.Split(input, "\n")
	parts := []string{""}
	processed := []string{}
	for i := range lines {
		line := ""
		for _, val := range lines[i] {
			newParts := []string{""}
			if val >= '0' && val <= '9' {
				line += string(val)
				parts = newParts
				continue
			}
			for j := range parts {
				parts[j] += string(val)
				for k, candidate := range candidates {
					if candidate == parts[j] {
						line += fmt.Sprintf("%v", k+1)
						break
					}
					length := len(parts[j])
					if len(candidate) > length && candidate[length-1] == parts[j][length-1] {
						newParts = append(newParts, parts[j])
					}
				}
			}
			parts = newParts
		}
		processed = append(processed, line)
	}
	input = strings.Join(processed, "\n")
	part1(input)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)

	part1(input)
	part2(input)
}
