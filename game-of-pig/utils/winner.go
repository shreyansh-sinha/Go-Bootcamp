package winner

import "math/rand"

func FindWinner(holdScore1 int, holdScore2 int) string {

	totalScore1 := 0
	totalScore2 := 0
	chance := 0

	for totalScore1 < 100 || totalScore2 < 100 {

		curScore := 0

		if chance == 0 {

			for curScore < holdScore1 {
				curValue := rand.Intn(6) + 1
				curScore += curValue
			}

			totalScore1 += curScore
			curScore = 0
		} else {

			for curScore < holdScore2 {
				curValue := rand.Intn(6)
				curScore += curValue
			}

			totalScore2 += curScore
			curScore = 0
		}

		if totalScore1 >= 100 {

			return "A"
		} else if totalScore2 >= 100 {

			return "B"
		}

		chance = 1 - chance
	}
	return ""
}
