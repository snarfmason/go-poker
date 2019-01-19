package poker

import "strings"

type Hand struct {
	Cards [5]Card
}

func (self Hand) String() string {
	var strs [5]string
	for i := 0; i < 5; i++ {
		strs[i] = self.Cards[i].String()
	}
	return strings.Join(strs[:], " ")
}

func (self Hand) RankString() string {
	return "Not Implemented"
}

func ParseHand(str string) Hand {
	cards := strings.Split(str, " ")

	if len(cards) != 5 {
		panic("Wrong Length")
	}

	var hand Hand
	for i := 0; i < 5; i++ {
		hand.Cards[i] = ParseCard(cards[i])
	}

	return hand
}
