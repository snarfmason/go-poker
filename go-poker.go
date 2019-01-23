package main

import (
	"bufio"
	"fmt"
	"go-poker/poker"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(play(line))
		line, err = reader.ReadString('\n')
	}
}

func play(line string) string {
	hands := hands(line)

	winner := winner(hands)
	winners := winners(winner, hands)

	result := fmt.Sprintf(
		"%v, Winner: %v, Rank: %v",
		strings.TrimRight(line, "\n"),
		strings.Join(winners, ", "),
		winner.RankString())

	return result
}

func hands(line string) []poker.Hand {
	strs := strings.Split(line, "|")
	n_hands := len(strs)
	hands := make([]poker.Hand, n_hands)
	for i := 0; i < n_hands; i++ {
		hands[i] = poker.ParseHand(strs[i])
	}
	return hands
}

func winner(hands []poker.Hand) poker.Hand {
	winner := hands[0]
	for i := 1; i < len(hands); i++ {
		winner = poker.BetterHand(winner, hands[i])
	}
	return winner
}

func winners(winner poker.Hand, hands []poker.Hand) []string {
	winners := make([]string, 0)
	for i := 0; i < len(hands); i++ {
		if poker.EqualHands(winner, hands[i]) {
			winners = append(winners, hands[i].String())
		}
	}
	return winners
}
