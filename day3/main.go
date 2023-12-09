package main

import (
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) {
	sum := 0
	for i := range lines {
		lines[i] += "."
	}
	for i := range lines {
		start := -1
		for j, c := range lines[i] {
			if c >= '0' && c <= '9' {
				if start == -1 {
					start = j
				}
			} else {
				if start != -1 {
				Loop:
					for x := start - 1; x <= j; x++ {
						for _, y := range []int{i, i + 1, i - 1} {
							if x < 0 || x >= len(lines[i]) || y < 0 || y >= len(lines)-1 {
								continue
							}
							char := lines[y][x]
							if char != '.' && !(char >= '0' && char <= '9') {
								num, _ := strconv.Atoi(lines[i][start:j])
								sum += num
								break Loop
							}
						}
					}
					start = -1
				}
			}
		}
	}
	println(sum)
}

func findNumberStart(line string, x int) int {
	for x >= 0 && line[x] >= '0' && line[x] <= '9' {
		x -= 1
	}

	return x + 1
}

func findNumbers(lines []string, x int, y int) []int {
	result := []int{}
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i > len(lines)-2 {
			continue
		}
		found := map[int]bool{}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j > len(lines[i])-1 {
				continue
			}
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				found[findNumberStart(lines[i], j)] = true
			}
		}
		for start := range found {
			end := start
			for end < len(lines[i]) && lines[i][end] >= '0' && lines[i][end] <= '9' {
				end += 1
			}
			num, _ := strconv.Atoi(lines[i][start:end])
			result = append(result, num)
		}
	}

	return result
}

func part2(lines []string) {
	total := 0
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '*' {
				nums := findNumbers(lines, x, y)
				if len(nums) == 2 {
					total += nums[0] * nums[1]
				}
			}
		}
	}

	println(total)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)

	part1(strings.Split(input, "\n"))
	part2(strings.Split(input, "\n"))
}
