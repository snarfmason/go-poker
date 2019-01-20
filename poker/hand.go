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
	switch {
	case self.isStraightFlush():
		return StraightFlush
	case self.isFourOfAKind():
		return FourOfAKind
	case self.isFullHouse():
		return FullHouse
	case self.isFlush():
		return Flush
	case self.isStraight():
		return Straight
	case self.isThreeOfAKind():
		return ThreeOfAKind
	case self.isTwoPair():
		return TwoPair
	case self.isPair():
		return Pair
	}
	return HighCard
}

func (self Hand) isPair() bool {
	return len(self.rankSets()) == 4
}

func (self Hand) isTwoPair() bool {
	return len(self.rankSets()) == 3 && self.mostCommonRankSize() == 2
}

func (self Hand) isThreeOfAKind() bool {
	return len(self.rankSets()) == 3 && self.mostCommonRankSize() == 3
}

func (self Hand) isStraight() bool {
	return self.isAllConsecutive() || self.isWraparoundStraight()
}

func (self Hand) isFlush() bool {
	return self.isAllSameSuit()
}

func (self Hand) isFullHouse() bool {
	return len(self.rankSets()) == 2 && self.mostCommonRankSize() == 3
}

func (self Hand) isFourOfAKind() bool {
	return len(self.rankSets()) == 2 && self.mostCommonRankSize() == 4
}

func (self Hand) isStraightFlush() bool {
	return self.isStraight() && self.isFlush()
}

func (self Hand) rankSets() []int {
	var ranks [5]int
	for i := 0; i < 5; i++ {
		ranks[i] = self.Cards[i].Value()
	}
	sort.Ints(ranks[:])

	// this part is ranks.uniq in Ruby
	rankSets := make([]int, 1)
	rankSets[0] = ranks[0]
	for i := 1; i < 5; i++ {
		if ranks[i] != ranks[i-1] {
			rankSets = append(rankSets, ranks[i])
		}
	}
	return rankSets
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

func (self Hand) mostCommonRankSize() int {
	max := 0
	for _, size := range self.rankSizes() {
		if size > max {
			max = size
		}
	}
	return max
}

func (self Hand) isWraparoundStraight() bool {
	ranks := self.rankSets()
	return len(ranks) == 5 && ranks[3] == 5 && ranks[4] == 14
}

func (self Hand) isAllConsecutive() bool {
	ranks := self.rankSets()
	return len(ranks) == 5 && ranks[4]-ranks[0] == 4
}

func (self Hand) isAllSameSuit() bool {
	theSuit := self.Cards[0].Suit
	for i := 1; i < 5; i++ {
		if self.Cards[i].Suit != theSuit {
			return false
		}
	}
	return true
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
	r1 := h1.Rank()
	r2 := h2.Rank()
	if r1 > r2 {
		return h1
	}
	if r2 > r1 {
		return h2
	}

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
	r1 := h1.Rank()
	r2 := h2.Rank()
	if r1 != r2 {
		return false
	}

	v1 := h1.Value()
	v2 := h2.Value()
	for i := 0; i < 9; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
