package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "queen", "jack", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)
	totVal := card1Val + card2Val
	switch {
	case totVal == 22:
		return "P"
	case totVal >= 17 && totVal <= 20:
		return "S"
	case totVal >= 12 && totVal <= 16:
		switch {
		case dealerCardVal >= 7:
			return "H"
		default:
			return "S"
		}
	case totVal == 21:
		switch dealerCardVal {
		case 11, 10:
			return "S"
		default:
			return "W"
		}
	default:
		return "H"
	}

}
