package poker

import (
	"fmt"
)

type Card struct {
	Rank byte
	Suit byte
}

func (self Card) String() string {
	return fmt.Sprintf("%c%c", self.Rank, self.Suit)
}

func (self Card) Value() int {
	return RankValue(self.Rank)
}

func ParseCard(str string) Card {
	return Card{str[0], str[1]}
}
