package problems

import (
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func FiftyFour() int {
	lines, _ := utils.GetFileLines("inputs/54.txt")

	wins := 0

	for _, line := range lines {
		cards := strings.Split(line, " ")

		player1 := utils.NewHand(cards[:5])
		player2 := utils.NewHand(cards[5:])

		p1Rank, p1Value, p1kickers := player1.Value()
		p2Rank, p2Value, p2kickers := player2.Value()

		if p1Rank > p2Rank {
			wins++
		} else if p1Rank == p2Rank {
			if p1Value > p2Value {
				wins++
			} else if p1Value == p2Value {
				for i := 0; i < len(p1kickers); i++ {
					k1 := p1kickers[i]
					k2 := p2kickers[i]

					if k1 > k2 {
						wins++
					} else if k1 == k2 {
						continue
					} else {
						break
					}
				}
			}
		}
	}

	return wins
}
