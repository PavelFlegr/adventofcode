package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func part1(lines []string) {
	instructions := lines[0]
	r, _ := regexp.Compile(`(.+) = \((.+), (.+)\)`)
	navigation := map[string][]string{}
	lines = lines[2:]
	for _, line := range lines {
		if line == "" {
			continue
		}

		match := r.FindStringSubmatch(line)
		navigation[match[1]] = append(navigation[match[1]], match[2], match[3])
	}

	current := "AAA"
	i := 0
	mod := len(instructions)
	for current != "ZZZ" {
		dir := 0
		if instructions[i%mod] == 'R' {
			dir = 1
		}
		current = navigation[current][dir]
		i += 1
	}

	fmt.Println(i)
}

func part2(lines []string) {
	instructions := lines[0]
	r, _ := regexp.Compile(`(.+) = \((.+), (.+)\)`)
	navigation := map[string][]string{}
	lines = lines[2:]
	for _, line := range lines {
		if line == "" {
			continue
		}

		match := r.FindStringSubmatch(line)
		navigation[match[1]] = append(navigation[match[1]], match[2], match[3])
	}

	nodes := []string{}
	for node := range navigation {
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	steps := []int{}
	mod := len(instructions)
	for _, current := range nodes {
		i := 0
		for current[2] != 'Z' {
			dir := 0
			if instructions[i%mod] == 'R' {
				dir = 1
			}
			current = navigation[current][dir]
			i += 1
		}
		steps = append(steps, i)
	}

	fmt.Println(LCM(steps[0], steps[1], steps[2:]...))
}

// Shamelessly copied LCM calculation from the interwebz
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")

	part1(lines)
	start := time.Now()
	part2(lines)
	fmt.Println(time.Since(start))
}
