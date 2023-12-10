package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type point struct {
	x int
	y int
}

func part1(lines []string) (map[point][]point, point) {
	edges := map[point][]point{}
	var start point
	for y := range lines {
		if lines[y] == "" {
			continue
		}
		for x, c := range lines[y] {
			var neighbours []point

			switch c {
			case '7':
				neighbours = []point{
					{x: x - 1, y: y},
					{x: x, y: y + 1},
				}
			case 'F':
				neighbours = []point{
					{x: x + 1, y: y},
					{x: x, y: y + 1},
				}
			case 'L':
				neighbours = []point{
					{x: x + 1, y: y},
					{x: x, y: y - 1},
				}
			case '|':
				neighbours = []point{
					{x: x, y: y - 1},
					{x: x, y: y + 1},
				}
			case 'J':
				neighbours = []point{
					{x: x - 1, y: y},
					{x: x, y: y - 1},
				}
			case '-':
				neighbours = []point{
					{x: x - 1, y: y},
					{x: x + 1, y: y},
				}
			case 'S':
				start = point{x: x, y: y}
			}

			edges[point{x: x, y: y}] = append(edges[point{x: x, y: y}], neighbours...)
		}
	}

	for src, dst := range edges {
		if slices.Contains(dst, start) {
			edges[start] = append(edges[start], src)
		}
	}

	prev := start
	current := edges[start][0]
	i := 1
	for {
		for _, e := range edges[current] {
			if e != prev {
				prev = current
				current = e
				break
			}

		}
		i += 1
		if current == start {
			break
		}
	}

	println(i / 2)
	return edges, start
}

func part2(lines []string, area [][]rune, edges map[point][]point, start point) {
	prev := start
	include := true
	current := edges[start][0]
	for {
		for _, e := range edges[current] {
			if e != prev {
				prev = current
				current = e
				break
			}

		}
		area[1+current.y*2][1+current.x*2] = rune(lines[current.y][current.x])
		if current.x != prev.x {
			area[1+current.y*2][(1 + current.x + prev.x)] = '-'
		} else {
			area[1+(current.y+prev.y)][1+current.x*2] = '|'
		}
		if current == start {
			if include {
				include = false
				continue
			}
			break
		}
	}

	queue := []point{{x: 0, y: 0}}
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			p := queue[i]
			if !(p.x >= 0 && p.y >= 0 && p.x < len(area[0]) && p.y < len(area)) {
				continue
			}
			c := area[p.y][p.x]
			if c == ' ' {
				area[p.y][p.x] = 'O'
				queue = append(queue, []point{{x: p.x - 1, y: p.y}, {x: p.x + 1, y: p.y}, {x: p.x, y: p.y - 1}, {x: p.x, y: p.y + 1}}...)
			}
		}
		queue = queue[count:]
	}

	count := 0
	for i := range lines {
		for j := range lines[i] {
			c := area[1+i*2][1+j*2]
			if c == ' ' {
				count += 1
			}
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	println(count)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")
	edges, start := part1(lines)
	rows := len(lines)*2 + 2
	cols := len(lines[0])*2 + 2
	area := make([][]rune, rows)
	for i := range area {
		area[i] = make([]rune, cols)
		for j := range area[i] {
			area[i][j] = ' '
		}
	}
	part2(lines, area, edges, start)
}
