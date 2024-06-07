package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"example.com/winner"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: ./Story2 <input1> <input2>")
		return
	}

	input1, err := parseInput(flag.Arg(0))
	if err != nil {
		fmt.Println("Error parsing input 1:", err)
		return
	}

	input2, err := parseInput(flag.Arg(1))
	if err != nil {
		fmt.Println("Error parsing input 2:", err)
		return
	}

	totalGames := 10 * (input2.End - 1)

	for score1 := input1.Start; score1 <= input1.End; score1++ {
		winCount := 0

		for score2 := input2.Start; score2 <= input2.End; score2++ {
			if score1 != score2 {

				for games := 0; games < 10; games++ {
					winner := winner.FindWinner(score1, score2)
					if winner == "A" {
						winCount++
					}
				}
			}
		}

		winPercentage := float32(winCount) * 100 / float32(totalGames)
		lossPercentage := 100.0 - winPercentage
		lossCount := totalGames - winCount

		fmt.Printf("Result: Wins, losses staying at k = %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
			score1, winCount, totalGames, winPercentage, lossCount, totalGames, lossPercentage)
	}
}

type Range struct {
	Start, End int
}

func parseInput(input string) (Range, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return Range{}, fmt.Errorf("invalid input format: %s", input)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, fmt.Errorf("invalid input format: %s", input)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, fmt.Errorf("invalid input format: %s", input)
	}

	return Range{Start: start, End: end}, nil
}
