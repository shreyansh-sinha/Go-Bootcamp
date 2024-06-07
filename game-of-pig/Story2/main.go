package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/winner"
)

func main() {

	args := os.Args[1:]

	holdScore1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid value")
		return
	}

	for holdScore2 := 0; holdScore2 <= 100; holdScore2++ {

		totalGames := 10
		winCount := 0

		for game := 0; game < totalGames; game++ {

			winner := winner.FindWinner(holdScore1, holdScore2)

			if winner == "A" {
				winCount++
			}
		}

		winPercentage := float32(winCount) * 100 / float32(totalGames)
		lossPercentage := 100.0 - winPercentage

		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
			holdScore1, holdScore2, winCount, totalGames, winPercentage, totalGames-winCount, totalGames, lossPercentage)
	}
}
