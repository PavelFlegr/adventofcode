package main

import (
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	src int
	dst int
	cnt int
}

func useMap(m []mapping, src int) int {
	for _, m := range m {
		if src >= m.src && src <= m.src+m.cnt {
			return src - m.src + m.dst
		}
	}

	return src
}

func part1(sections []string) {
	seeds := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	maps := map[string][]mapping{}
	sections = sections[1:]
	for s := range sections {
		lines := strings.Split(sections[s], "\n")
		section := strings.Split(strings.Split(lines[0], "-")[2], " ")[0]
		lines = lines[1:]
		for l := range lines {
			if lines[l] == "" {
				continue
			}
			parts := strings.Split(lines[l], " ")
			dst, _ := strconv.Atoi(parts[0])
			src, _ := strconv.Atoi(parts[1])
			cnt, _ := strconv.Atoi(parts[2])
			maps[section] = append(maps[section], mapping{src: src, dst: dst, cnt: cnt})
		}
	}

	locations := []int{}

	for s := range seeds {
		if s%2 == 1 {
			continue
		}
		seed, _ := strconv.Atoi(seeds[s])
		soil := useMap(maps["soil"], seed)
		fertilizer := useMap(maps["fertilizer"], soil)
		water := useMap(maps["water"], fertilizer)
		light := useMap(maps["light"], water)
		temperature := useMap(maps["temperature"], light)
		humidity := useMap(maps["humidity"], temperature)
		location := useMap(maps["location"], humidity)
		locations = append(locations, location)
	}

	min := locations[0]
	for l := range locations {
		loc := locations[l]
		if loc < min {
			min = loc
		}
	}

	println(min)
}

type Range struct {
	start int
	end   int
}

func mapRange(m []mapping, r Range) Range {
	count := r.end - r.start
	r.start = useMap(m, r.start)
	r.end = r.start + count

	return r
}

func part2(sections []string) {
	seedss := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	currentRanges := []Range{}

	for i := range seedss {
		if i%2 == 1 {
			continue
		}
		start, _ := strconv.Atoi(seedss[i])
		count, _ := strconv.Atoi(seedss[i+1])
		currentRanges = append(currentRanges, Range{start: start, end: start + count + 1})
	}
	maps := map[string][]mapping{}
	sections = sections[1:]
	for s := range sections {
		lines := strings.Split(sections[s], "\n")
		section := strings.Split(strings.Split(lines[0], "-")[2], " ")[0]
		lines = lines[1:]
		for l := range lines {
			if lines[l] == "" {
				continue
			}
			parts := strings.Split(lines[l], " ")
			dst, _ := strconv.Atoi(parts[0])
			src, _ := strconv.Atoi(parts[1])
			cnt, _ := strconv.Atoi(parts[2])
			maps[section] = append(maps[section], mapping{src: src, dst: dst, cnt: cnt})
		}
	}
	var newRange Range

	for _, i := range []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"} {
		ranges := []Range{}
		for _, s := range currentRanges {
			newRange.start = s.start
			for _, m := range maps[i] {
				mEnd := m.src + m.cnt
				if mEnd > s.start && mEnd < s.end {
					newRange.end = mEnd
					ranges = append(ranges, mapRange(maps[i], newRange))
					newRange.start = mEnd + 1
				} else if m.src > newRange.start && m.src < s.end {
					newRange.end = m.src - 1
					ranges = append(ranges, mapRange(maps[i], newRange))
					newRange.start = m.src
					if mEnd < s.end {
						newRange.end = mEnd
						ranges = append(ranges, mapRange(maps[i], newRange))
						newRange.start = mEnd + 1
					}
				}
			}

			newRange.end = s.end
			ranges = append(ranges, mapRange(maps[i], newRange))
		}
		currentRanges = ranges
	}

	min := currentRanges[0].start
	for l := range currentRanges {
		loc := currentRanges[l]
		if loc.start < min {
			min = loc.start
		}
	}

	println(min)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)

	part1(strings.Split(input, "\n\n"))
	part2(strings.Split(input, "\n\n"))
}
