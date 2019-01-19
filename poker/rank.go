package poker

func RankValue(cardRank byte) int {
	switch cardRank {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	// '2' through '9'
	return int(cardRank - 48)
}
