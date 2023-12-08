package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	_ "embed"

	"github.com/sirupsen/logrus"
)

//go:embed input.txt
var inputB []byte

//go:embed test.txt
var testB []byte

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards string
	bet   int
	typ   Type
}

var cardsOrder = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func (h *Hand) IsStrongerThan(otherH *Hand) bool {
	if h.typ > otherH.typ {
		return true
	}

	if h.typ < otherH.typ {
		return false
	}

	// lenghts are considered equal
	for i := range h.cards {
		hIdx := slices.Index(cardsOrder, string(h.cards[i]))
		otherIdx := slices.Index(cardsOrder, string(otherH.cards[i]))

		if hIdx > otherIdx {
			return true
		}

		if otherIdx > hIdx {
			return false
		}
	}

	logrus.Errorf("equal hands `%s` and `%s`, returning true", h.cards, otherH.cards)
	return true
}

func main() {
	input := string(inputB)
	lines := strings.Split(input, "\n")

	hands := make([]*Hand, len(lines))
	for i, l := range lines {
		hand, err := parseHand(l)
		if err != nil {
			logrus.WithError(err).Error("failed to parse hand")
			continue
		}

		hands[i] = hand
	}

	earnings := 0
	for _, h := range hands {
		rank := findRank(hands, h)
		fmt.Println("cards:", h.cards, "bet:", h.bet, "rank:", rank)
		earnings = earnings + rank*h.bet
	}

	fmt.Println(earnings)
}

func parseHand(line string) (*Hand, error) {
	splitted := strings.Split(line, " ")
	cards := splitted[0]
	bet, err := strconv.Atoi(splitted[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert bet string to int: %w", err)
	}

	return &Hand{
		cards: cards,
		bet:   bet,
		typ:   findCardsType(cards),
	}, nil
}

func findCardsType(cards string) Type {
	cardsM := make(map[rune]int)
	for _, c := range cards {
		cardsM[c] = cardsM[c] + 1
	}

	var twos int
	for _, count := range cardsM {
		if count == 5 {
			return FiveOfAKind
		}

		if count == 4 {
			return FourOfAKind
		}

		if count == 3 {
			return ThreeOfAKind
		}

		if count == 2 {
			twos++
		}
	}

	if twos == 2 {
		return TwoPairs
	}

	if twos == 1 {
		return OnePair
	}

	return HighCard
}

func findRank(hands []*Hand, refH *Hand) int {
	rank := 1 // base starts at 1
	for _, h := range hands {
		if refH == h {
			continue
		}

		if refH.IsStrongerThan(h) {
			rank++
		}
	}

	return rank
}
