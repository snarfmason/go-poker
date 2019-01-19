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
		hands := hands(line)

		winners := winners(hands)

		output := fmt.Sprintf(
			"%v, Winner: %v, Rank: %v",
			strings.TrimRight(line, "\n"),
			strings.Join(winners, ", "),
			"none") // winners[0].RankString()

		fmt.Println(output)
		line, err = reader.ReadString('\n')
	}
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

func winners(hands []poker.Hand) []string {
	winners := make([]string, 0)
	// just return the first hand for now
	return append(winners, hands[0].String())
}
