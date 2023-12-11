package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

type point struct {
	x int
	y int
}

func part1(lines []string) {
	galaxies := []point{}
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				galaxies = append(galaxies, point{x: j, y: i})
			}
		}
	}

	updatedGalaxies := slices.Clone(galaxies)
	for i := range lines {
		hasGalaxy := false
		j := 0
		for j = range lines[i] {
			if lines[i][j] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for k := range galaxies {
				if galaxies[k].y > i {
					updatedGalaxies[k].y += 1
				}
			}
		}
	}

	for i := range lines[0] {
		hasGalaxy := false
		j := 0
		for j = range lines {
			if lines[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for k := range galaxies {
				if galaxies[k].x > i {
					updatedGalaxies[k].x += 1
				}
			}
		}
	}

	total := 0.0
	for i := range updatedGalaxies {
		for j := 0; j < i; j++ {
			xDistance := math.Abs(float64(updatedGalaxies[i].x - updatedGalaxies[j].x))
			yDistance := math.Abs(float64(updatedGalaxies[i].y - updatedGalaxies[j].y))
			distance := xDistance + yDistance
			total += distance
		}
	}

	println(int(total))
}

func part2(lines []string) {
	galaxies := []point{}
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				galaxies = append(galaxies, point{x: j, y: i})
			}
		}
	}

	updatedGalaxies := slices.Clone(galaxies)
	for i := range lines {
		hasGalaxy := false
		j := 0
		for j = range lines[i] {
			if lines[i][j] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for k := range galaxies {
				if galaxies[k].y > i {
					updatedGalaxies[k].y += 1e6 - 1
				}
			}
		}
	}

	for i := range lines[0] {
		hasGalaxy := false
		j := 0
		for j = range lines {
			if lines[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for k := range galaxies {
				if galaxies[k].x > i {
					updatedGalaxies[k].x += 1e6 - 1
				}
			}
		}
	}

	total := 0.0
	for i := range updatedGalaxies {
		for j := 0; j < i; j++ {
			xDistance := math.Abs(float64(updatedGalaxies[i].x - updatedGalaxies[j].x))
			yDistance := math.Abs(float64(updatedGalaxies[i].y - updatedGalaxies[j].y))
			distance := xDistance + yDistance
			total += distance
		}
	}

	println(int(total))
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]
	start := time.Now()
	fmt.Println(time.Since(start))
	fmt.Println(time.Since(start))
}
