package main

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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var sum int = ParseCard(card1) + ParseCard(card2)
	var dealerSum int = ParseCard(dealerCard)
	if sum == 22 { // Both cards are aces.
		return "P"
	} else if sum == 21 {
		switch dealerCard {
		case "ace", "ten", "queen":
			return "S"
		default:
			return "W"
		}
	} else if sum >= 17 && sum <= 20 {
		return "S"
	} else if sum >= 12 && sum <= 16 {
		if dealerSum >= 7 {
			return "H"
		} else {
			return "S"
		}
	} else if sum <= 11 {
		return "H"
	}
	return "S"
}

func main() {
	value := ParseCard("ace")
	println("Parse card (expected - 11):", value)
	println("First turn (expected - P):", FirstTurn("ace", "ace", "jack"))
	println("First turn (expected - S):", FirstTurn("ace", "king", "ace"))
	println("First turn (expected - H):", FirstTurn("five", "queen", "ace"))
	println("First turn (expected - S):", FirstTurn("king", "ace", "queen"))
}
