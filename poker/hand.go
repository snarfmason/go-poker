package poker

import (
	"sort"
	"strings"
)

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
	return RankString(self.Rank())
}

func (self Hand) Value() [9]int {
	if self.isWraparoundStraight() {
		return [9]int{0, 0, 0, 0, 5, 4, 3, 2, 1}
	}

	var fourRank int
	var threeRank int
	pairs := make([]int, 0)
	singles := make([]int, 0)

	for key, val := range self.rankSizes() {
		if val == 4 {
			fourRank = key
		} else if val == 3 {
			threeRank = key
		} else if val == 2 {
			pairs = append(pairs, key)
		} else {
			singles = append(singles, key)
		}
	}
	sort.Ints(pairs)
	sort.Ints(singles)
	for i := len(pairs); i < 2; i++ {
		pairs = append(pairs, 0)
	}
	for i := len(singles); i < 5; i++ {
		singles = append(singles, 0)
	}

	return [9]int{
		fourRank,
		threeRank,
		pairs[1],
		pairs[0],
		singles[4],
		singles[3],
		singles[2],
		singles[1],
		singles[0]}
}

func (self Hand) Rank() int {
	return -1
}

func (self Hand) rankSizes() map[int]int {
	frequencies := make(map[int]int)
	var c Card
	for i := 0; i < 5; i++ {
		c = self.Cards[i]
		frequencies[c.Value()]++
	}
	return frequencies
}

func (self Hand) isWraparoundStraight() bool {
	return false
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

func BetterHand(h1, h2 Hand) Hand {
	v1 := h1.Value()
	v2 := h2.Value()
	for i := 0; i < 9; i++ {
		if v1[i] > v2[i] {
			return h1
		}
		if v2[i] > v1[i] {
			return h2
		}
	}
	return h1 // hands are equal if we got here
}

func EqualHands(h1, h2 Hand) bool {
	v1 := h1.Value()
	v2 := h2.Value()
	for i := 0; i < 9; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
