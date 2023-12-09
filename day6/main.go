package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func part1(lines []string) {
	races := []Race{}
	current := ""
	for _, c := range strings.Split(lines[0], ":")[1] + " " {
		if c >= '0' && c <= '9' {
			current += string(c)
		} else {
			if current != "" {
				parsed, _ := strconv.Atoi(current)
				races = append(races, Race{time: parsed})
				current = ""
			}
			current = ""
		}
	}
	i := 0
	for _, c := range strings.Split(lines[1], ":")[1] + " " {
		if c >= '0' && c <= '9' {
			current += string(c)
		} else {
			if current != "" {
				parsed, _ := strconv.Atoi(current)
				races[i].distance = parsed
				i += 1
			}
			current = ""
		}
	}

	prod := 1

	for _, race := range races {
		target := race.time
		dis := math.Sqrt(float64((target * target) - 4*race.distance))
		h1 := (float64(target) - dis) / 2
		h2 := (float64(target) + dis) / 2
		prod = prod * int((math.Ceil(h2) - math.Floor(h1) - 1))
	}

	println(prod)
}

func part2(lines []string) {
	targetS := ""
	for _, c := range strings.Split(lines[0], ":")[1] {
		if c >= '0' && c <= '9' {
			targetS += string(c)
		}
	}
	distanceS := ""
	for _, c := range strings.Split(lines[1], ":")[1] {
		if c >= '0' && c <= '9' {
			distanceS += string(c)
		}
	}
	target, _ := strconv.Atoi(targetS)
	distance, _ := strconv.Atoi(distanceS)
	dis := math.Sqrt(float64((target * target) - 4*distance))
	h1 := (float64(target) - dis) / 2
	h2 := (float64(target) + dis) / 2
	println(int((math.Ceil(h2) - math.Floor(h1) - 1)))
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}
