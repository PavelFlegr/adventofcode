package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	counts map[rune]int
	cards  string
	bid    int
	score  int
}

func part1(lines []string) {
	hands := []Hand{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		hand := Hand{}
		hand.counts = map[rune]int{}
		hand.cards = strings.Split(line, " ")[0]
		hand.bid, _ = strconv.Atoi(strings.Split(line, " ")[1])

		for _, c := range hand.cards {
			hand.counts[c] += 1
		}

		for _, c := range hand.counts {
			if c == 5 {
				hand.score = 6
				break
			} else if c == 4 {
				hand.score = 5
				break
			} else if c == 3 {
				if hand.score == 1 {
					hand.score = 4
					break
				}
				hand.score = 3
			} else if c == 2 {
				if hand.score == 3 {
					hand.score = 4
					break
				}
				if hand.score == 1 {
					hand.score = 2
					break
				}
				hand.score = 1
			}
		}

		hands = append(hands, hand)
	}

	scores := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		tmp := cmp.Compare(a.score, b.score)
		if tmp != 0 {
			return tmp
		}
		for i := range a.cards {
			scoreA := string(a.cards[i])
			scoreB := string(b.cards[i])
			x := scores[scoreA]
			y := scores[scoreB]
			res := cmp.Compare(x, y)
			if res != 0 {
				return res
			}
		}
		return 0
	})

	winnings := 0
	for i, hand := range hands {
		winnings += ((i + 1) * hand.bid)
	}
	fmt.Println(winnings)
}

func part2(lines []string) {
	scores := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 1,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	hands := []Hand{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		hand := Hand{}
		hand.counts = map[rune]int{}
		hand.cards = strings.Split(line, " ")[0]
		hand.bid, _ = strconv.Atoi(strings.Split(line, " ")[1])

		for _, c := range hand.cards {
			hand.counts[c] += 1
		}

		max := 0
		maxI := '0'
		for i, c := range hand.counts {
			if i == 'J' {
				continue
			}
			if c > max || c == max && scores[string(i)] > scores[string(maxI)] {
				max = c
				maxI = i
			}
		}
		if max == 0 {
			hand.counts = map[rune]int{'A': 5}
		} else {
			hand.counts[maxI] += hand.counts['J']
		}

		for i, c := range hand.counts {
			if i == 'J' {
				continue
			}
			if c == 5 {
				hand.score = 6
				break
			} else if c == 4 {
				hand.score = 5
				break
			} else if c == 3 {
				if hand.score == 1 {
					hand.score = 4
					break
				}
				hand.score = 3
			} else if c == 2 {
				if hand.score == 3 {
					hand.score = 4
					break
				}
				if hand.score == 1 {
					hand.score = 2
					break
				}
				hand.score = 1
			}
		}

		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		tmp := cmp.Compare(a.score, b.score)
		if tmp != 0 {
			return tmp
		}
		for i := range a.cards {
			scoreA := string(a.cards[i])
			scoreB := string(b.cards[i])
			x := scores[scoreA]
			y := scores[scoreB]
			res := cmp.Compare(x, y)
			if res != 0 {
				return res
			}
		}
		return 0
	})

	winnings := 0
	for i, hand := range hands {
		winnings += ((i + 1) * hand.bid)
	}
	fmt.Println(winnings)
}

func main() {
	buf, _ := os.ReadFile("input")
	input := string(buf)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}
